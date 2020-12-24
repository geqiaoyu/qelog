package qezap

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/huzhongqing/qelog/qezap/push"
)

type WriteRemoteConfig struct {
	Transport  string // 支持 http || grpc 默认grpc
	Addrs      []string
	ModuleName string

	MaxConcurrent int           // 默认 1 个并发
	MaxPacket     int           // 默认不缓冲
	WriteTimeout  time.Duration // 默认不超时
}

func (cfg WriteRemoteConfig) Validate() error {
	if len(cfg.Addrs) == 0 {
		return fmt.Errorf("address required, grpc [ip:port]  http[url]")
	}
	if cfg.ModuleName == "" {
		return fmt.Errorf("moduleName required")
	}
	return nil
}

func NewWriteRemoteConfig(addrs []string, moduleName string) WriteRemoteConfig {
	return WriteRemoteConfig{
		Transport:     "grpc",
		Addrs:         addrs,
		ModuleName:    moduleName,
		MaxConcurrent: 50,
		// 包的大小对写入效率有着比较重要的影响。 当设置 1MB时，会快于 64KB
		// 但是小对象对于GC相对更加友好
		// 63KB 未选择 64KB，留1kb给GRPC本身携带信息。 (grpc 默认最大4MB一个包)
		MaxPacket:    63 << 10,
		WriteTimeout: 5 * time.Second,
	}
	// 如果超出并发限制，直接写入文件，缓慢背景发送
}

type WriteRemote struct {
	cfg    WriteRemoteConfig
	pusher push.Pusher

	packets *Packets

	once sync.Once
}

func NewWriteRemote(cfg WriteRemoteConfig) *WriteRemote {
	if err := cfg.Validate(); err != nil {
		panic("config validate error " + err.Error())
	}
	wr := &WriteRemote{
		cfg:     cfg,
		packets: NewPackets(cfg.MaxPacket),
	}

	wr.once.Do(func() {
		go wr.initPusher()
		go wr.backgroundRetry()
		go wr.pullPacket()
	})
	return wr
}

func (wr *WriteRemote) Write(b []byte) (n int, err error) {
	data, flush := wr.packets.AddPacket(b)
	if flush {
		wr.push(push.NewPacket(wr.cfg.ModuleName, data.val))
		wr.packets.Free(data)
	}
	return len(b), nil
}

func (wr *WriteRemote) pullPacket() {
	tick := time.NewTicker(2 * time.Second)
	for {
		select {
		case <-tick.C:
			data, flush := wr.packets.PullPacket()
			if flush {
				wr.push(push.NewPacket(wr.cfg.ModuleName, data.Val()))
				wr.packets.Free(data)
			}
		}
	}
}

func (wr *WriteRemote) push(in *push.Packet) {
	if wr.pusher == nil {
		_, _ = wr.packets.WriteBakPacket(in)
		return
	}
	// 如果发送者满负荷，则直接丢文件
	if wr.pusher.Concurrent() >= wr.cfg.MaxConcurrent {
		_, _ = wr.packets.WriteBakPacket(in)
		return
	}

	ctx := context.Background()
	if wr.cfg.WriteTimeout > 0 {
		ctx, _ = context.WithTimeout(context.Background(), wr.cfg.WriteTimeout)
	}

	go func() {
		if err := wr.pusher.PushPacket(ctx, in); err != nil {
			if err == push.ErrUnavailable {
				// 只有当服务不可用时，放入错误备份文件里
				_, _ = wr.packets.WriteBakPacket(in)
			}
			log.Printf("write remote push packet %s\n", err.Error())
		}
	}()
}

func (wr *WriteRemote) initPusher() {
	// 在发送的时候，才去链接， 如果链接不通，不能影响主进程
	tick := time.NewTicker(time.Second)
	for {
		if wr.pusher == nil {
			if wr.cfg.Transport == "http" {
				pusher, err := push.NewHttpPush(wr.cfg.Addrs[0], wr.cfg.MaxConcurrent)
				if err != nil {
					log.Printf("init http push error %s\n", err.Error())
					goto next
				}
				wr.pusher = pusher
			} else {
				pusher, err := push.NewGRPCPush(wr.cfg.Addrs, wr.cfg.MaxConcurrent)
				if err != nil {
					log.Printf("init grpc push error %s\n", err.Error())
					goto next
				}
				wr.pusher = pusher
			}
			tick.Stop()
			return
		}
	next:
		select {
		case <-tick.C:
		}
	}
}

func (wr *WriteRemote) backgroundRetry() {
	tick := time.NewTicker(200 * time.Millisecond)
	for {
		select {
		case <-tick.C:
			v := &push.Packet{}
			ok, err := wr.packets.ReadBakPacket(v)
			if err != nil {
				fmt.Println("packets retry", err.Error())
			}
			if ok {
				wr.push(v)
			}
		}
	}
}

func (wr *WriteRemote) Sync() error {
	data, flush := wr.packets.PullPacket()
	if flush {
		wr.push(push.NewPacket(wr.cfg.ModuleName, data.Val()))
		wr.packets.Free(data)
	}

	exit := make(chan struct{}, 1)
	go func() {
		for {
			if wr.pusher != nil && wr.pusher.Concurrent() == 0 {
				exit <- struct{}{}
				return
			}
			time.Sleep(50 * time.Millisecond)
		}
	}()

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	select {
	case <-ctx.Done():
		log.Println("sync ", ctx.Err())
	case <-exit:
	}

	return wr.packets.Close()
}