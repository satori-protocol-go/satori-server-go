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
		panic(err)
	}
	// 加载适配器插件
	instances, err := plugin.Load(ctx, nil)
	for _, instance := range instances {
		log.Info("process instance%:v", instance)
		// todo 将插件注入到适配器中心
		err := adaptor.Register(nil)
		if err != nil {
			cancel()
			panic(err)
		}
	}
	// 启动服务
	err = server.Start(ctx, nil)
	if err != nil {
		cancel()
		panic(err)
	}

}
