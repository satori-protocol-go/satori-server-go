package main

import (
	"context"

	"github.com/dezhishen/satori-server-go/pkg/adaptor"
	"github.com/dezhishen/satori-server-go/pkg/config"
	"github.com/dezhishen/satori-server-go/pkg/plugin"
	"github.com/dezhishen/satori-server-go/pkg/server"
	log "github.com/sirupsen/logrus"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	// 加载配置
	_, err := config.Load("./config.yaml")
	if err != nil {
		log.Panicf("load config file has err: %v", err)
	}
	// 加载适配器插件
	instances, err := plugin.Load(ctx, nil)
	if err != nil {
		log.Panicf("load plugin has err: %v", err)
	}
	adaptorHandler := adaptor.NewGRPCAdaptorInstance()
	for _, instance := range instances {
		log.Infof("start to process instance: %v", instance)
		// todo 将插件注入到适配器中心
		err := adaptorHandler.Register(nil)
		if err != nil {
			cancel()
			log.Panicf("register adaptor has err: %v", err)
		}
	}
	// 启动服务
	err = server.Start(ctx, nil)
	if err != nil {
		cancel()
		log.Panicf("start http server has err: %v", err)
	}

}
