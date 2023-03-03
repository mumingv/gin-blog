package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/mumingv/gin-blog/logger"
	"github.com/mumingv/gin-blog/settings"
)

func helloHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"mode": settings.Conf.Mode,
		"host": settings.Conf.MySQLConfig.Host,
	})
}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 中间件注册
	r.Use(logger.GinLogger())

	// 函数注册
	r.GET("/hello", helloHandler)

	return r
}
