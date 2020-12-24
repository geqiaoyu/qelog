package main

import (
	"strconv"
	"testing"
	"time"

	"github.com/huzhongqing/qelog/qezap"
	"go.uber.org/zap"
)

func TestQezapLoopWrite(t *testing.T) {

	// 持续写一段时间
	addrs := []string{"127.0.0.1:31082"}
	cfg := qezap.NewConfig(addrs, "example")
	cfg.WriteRemote.MaxConcurrent = 100
	// 如果设置 false，可以 addrs = nil
	//cfg.SetEnableRemote(false)
	// 如果对默认配置不满足，可直接设置
	qeLog := qezap.New(cfg, zap.DebugLevel)
	time.Sleep(2 * time.Second)
	go func() {
		for i := 0; i < 1000000; i++ {
			qeLog.Info(strconv.Itoa(i), zap.Int("index", i))
		}
	}()
	go func() {
		for i := 0; i < 1000000; i++ {
			qeLog.Warn(strconv.Itoa(i), zap.Int("index", i))
		}
	}()
	go func() {
		for i := 0; i < 1000000; i++ {
			qeLog.Error(strconv.Itoa(i), zap.Int("index", i))
		}
	}()
	time.Sleep(5 * time.Minute)
	qeLog.Sync()
}