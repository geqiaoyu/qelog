package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/huzhongqing/qelog/libs/logs"

	"github.com/huzhongqing/qelog/pkg/receiver"

	"github.com/huzhongqing/qelog/libs/mongo"

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
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	database, err := mongo.NewDatabase(ctx, cfg.MongoDB.URI,
		cfg.MongoDB.DataBase)
	if err != nil {
		log.Fatalln("mongo connect failed ", err.Error())
	}

	config.SetGlobalConfig(cfg)

	logs.InitQezap(nil, "")

	httpSrv := receiver.NewHTTPService(database)

	go func() {
		log.Println("http server listen ", cfg.ReceiverAddr)
		if err := httpSrv.Run(cfg.ReceiverAddr); err != nil {
			log.Fatalln("http server listen failed ", err.Error())
		}
	}()

	grpcSrv := receiver.NewGRPCService(database)
	go func() {
		log.Println("grpc server listen ", cfg.ReceiverGRPCAddr)
		if err := grpcSrv.Run(cfg.ReceiverGRPCAddr); err != nil {
			log.Fatalln("grpc server listen failed ", err.Error())
		}
	}()

	signalAccept()
	_ = httpSrv.Close()
	_ = grpcSrv.Close()
	_ = database.Client().Disconnect(nil)
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
