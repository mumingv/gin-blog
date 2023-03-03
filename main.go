package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github/com/mumingv/gin-blog/settings"
)

func main() {
	if err := settings.Init(); err != nil {
		fmt.Printf("Loading config failed, err: %v", err)
	}

	r := gin.Default()
	r.GET("/hello", func(ctx *gin.Context) {
		fmt.Println(settings.Conf)
		fmt.Println(settings.Conf.Mode)
		fmt.Println(settings.Conf.MySQLConfig.Host)
		ctx.JSON(200, gin.H{
			"mode": settings.Conf.Mode,
			"host": settings.Conf.MySQLConfig.Host,
		})
	})
	r.Run()
}
