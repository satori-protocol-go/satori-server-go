package cmd

import (
	"context"

	"github.com/dezhishen/satori-application-go/pkg/adaptor"
	"github.com/dezhishen/satori-application-go/pkg/application"
	"github.com/dezhishen/satori-application-go/pkg/config"
	"github.com/dezhishen/satori-application-go/pkg/plugin"
	log "github.com/sirupsen/logrus"
)

func Start() {
	ctx := context.Background()
	// 加载配置
	cfg, err := config.Load("./config.yaml")
	if err != nil {
		log.Panicf("load config file has err: %v", err)
	}
	// 构建服务对象
	appInstance, err := application.NewInstance(ctx, cfg)
	if err != nil {
		log.Panicf("start http server has err: %v", err)
	}
	// 加载适配器插件
	plugins, err := plugin.Load(ctx, cfg)
	if err != nil {
		log.Panicf("load plugin has err: %v", err)
	}
	adaptorRegistry := adaptor.NewSimpleRegistry(cfg)
	for _, e := range plugins {
		log.Infof("start to process instance: %v", e)
		switch e.Type() {
		case plugin.PLUGIN_ADAPTOR:
			err := adaptorRegistry.Register(e.Adatpor())
			if err != nil {
				log.Errorf("register adaptor has err: %v", err)
			}
			continue
		default:
			log.Warnf("unsupport type:%s", e.Type())
		}
	}
	err = appInstance.Start()
	if err != nil {
		log.Panicf("start http server has err: %v", err)
	}
}
