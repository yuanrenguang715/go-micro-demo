package main

import (
	"context"
	"go-micro-demo/server/handler"
	"go-micro-demo/server/proto/message"
	"log"
	"time"
	"utils/config"
	"utils/logger"

	grpc "github.com/micro/go-grpc"
	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro/server"
)

func main() {
	config.SrvCfgInit()
	logger.Init()

	service := grpc.NewService(
		micro.Name(config.SrvCfg.ServerName),
		micro.Registry(
			consul.NewRegistry(
				consul.Config(config.SrvCfg.ConsulConf),
			),
		),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*15),
		micro.WrapHandler(wrap),
		grpc.WithTLS(config.SrvCfg.TlsConf),
	)

	// optionally setup command line usage
	service.Init()

	// Register Handlers
	message.RegisterMessSrvHandler(service.Server(), new(handler.MessSrv))

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func wrap(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		start := time.Now()
		err := fn(ctx, req, rsp)

		//日志跟踪
		if config.SrvCfg.LogTrack {
			logger.Infof("%16s  %s", time.Since(start), req.Method())
		}

		return err
	}
}
