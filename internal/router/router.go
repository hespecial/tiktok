package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"tiktok/global"
	"tiktok/internal/router/api"
)

func StartServer() {
	gin.SetMode(global.Conf.App.Mode)
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	apiGroup := router.Group("/api")
	api.RegisterUserRouter(apiGroup)
	api.RegisterFollowRouter(apiGroup)

	log.Fatal(router.Run(fmt.Sprintf(":%s", global.Conf.App.Port)))
}
