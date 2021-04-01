package main

import (
	"flag"
	"fmt"
	"github.com/huzhongqing/qelog/infra/mongo"
	"github.com/huzhongqing/qelog/pkg/httpserver"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"

	"github.com/huzhongqing/qelog/infra/logs"

	"github.com/huzhongqing/qelog/pkg/common/model"
	"github.com/huzhongqing/qelog/pkg/receiver"

	"github.com/huzhongqing/qelog/pkg/config"
)

var (
	goVersion = ""
	buildTime = ""
	gitHash   = ""

	configPath = "./configs/config.toml"
	version    = false
)

func main() {
	flag.StringVar(&configPath, "f", "./configs/config.toml", "config file default(./configs/config.toml)")
	flag.BoolVar(&version, "v", false, "show version")
	flag.Parse()

	if version {
		fmt.Printf("goVersion: %s \nbuildTime: %s \ngitHash: %s \n", goVersion, buildTime, gitHash)
		return
	}

	cfg := config.InitConfig(configPath)
	if err := cfg.Validate(); err != nil {
		panic(fmt.Sprintf("config validate %s", err.Error()))
		return
	}

	config.SetGlobalConfig(cfg)

	logs.InitQezap(cfg.Logging.Addr, cfg.Logging.Module, cfg.Logging.Filename)

	slots := make([]mongo.ShardSlotConfig, 0)
	for _, v := range cfg.Sharding {
		slots = append(slots, mongo.ShardSlotConfig{
			Index:    v.Index,
			DataBase: v.DataBase,
			URI:      v.URI,
		})
	}
	sharding, err := mongo.NewSharding(mongo.MainConfig{
		DataBase: cfg.Main.DataBase,
		URI:      cfg.Main.URI,
	}, slots)
	if err != nil {
		logs.Qezap.Fatal("mongo connect failed ", zap.Error(err))
	}

	if err := model.SetGlobalShardingDB(sharding); err != nil {
		logs.Qezap.Fatal("SetGlobalShardingDB", zap.Error(err))
	}

	mainDB, err := sharding.MainDB()
	if err != nil {
		logs.Qezap.Fatal("mongo connect failed ", zap.Error(err))
	}
	if err := model.SetGlobalMainDB(mainDB); err != nil {
		logs.Qezap.Fatal("SetGlobalMainDB", zap.Error(err))
	}

	httpSrv := httpserver.NewHTTPServer(cfg.Env)
	receiver.RegisterRouter(httpSrv.Engine())
	go func() {
		fmt.Println("http server listen", cfg.ReceiverAddr)
		if err := httpSrv.Run(cfg.ReceiverAddr); err != nil {
			logs.Qezap.Fatal("http server listen failed", zap.Error(err))
		}
	}()

	grpcSrv := receiver.NewGRPCService()
	go func() {
		fmt.Println("gRPC server listen", cfg.ReceiverGRPCAddr)
		if err := grpcSrv.Run(cfg.ReceiverGRPCAddr); err != nil {
			logs.Qezap.Fatal("gRPC server listen failed", zap.Error(err))
		}
	}()

	signalAccept()
	_ = httpSrv.Close()
	_ = grpcSrv.Close()
	_ = sharding.Disconnect
}

func signalAccept() {
	// 不同的信号量不同的处理方式
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		si := <-ch
		switch si {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
