package cli

import (
	"go-micro-demo/server/proto/message"
	"utils/config"

	"github.com/micro/go-grpc"
	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/registry/consul"
)

var (
	MessClient message.MessSrvService

	// other client ...
)

func Init() {
	//default name
	messSrvName, messCli := initClient("")
	MessClient = message.NewMessSrvService(messSrvName, messCli)

	// other client ...
	//initClient("account")
	//initClient("trade")
}

func initClient(section string) (string, client.Client) {
	config.CliCfgInit(section)

	service := grpc.NewService(
		micro.Registry(
			consul.NewRegistry(
				consul.Config(config.CliCfg.ConsulConf),
			),
		),
		grpc.WithTLS(config.CliCfg.TlsConf),
	)
	service.Init()
	return config.CliCfg.ServerName, service.Client()
}
