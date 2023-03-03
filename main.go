package main

import (
	"fmt"
	"github.com/mumingv/gin-blog/logger"
	"github.com/mumingv/gin-blog/routers"
	"github.com/mumingv/gin-blog/settings"
)

func main() {
	// 加载配置
	if err := settings.Init(); err != nil {
		fmt.Printf("Loading config failed, err: %v\n", err)
	}

	// 初始化日志
	if err := logger.Init(settings.Conf.LogConfig, settings.Conf.Mode); err != nil {
		fmt.Printf("init logger failed, err: %v\n", err)
		return
	}

	// 注册路由
	r := routers.SetupRouter()
	if err := r.Run(fmt.Sprintf(":%d", settings.Conf.Port)); err != nil {
		fmt.Printf("server startup failed, err: %v\n", err)
	}
}
