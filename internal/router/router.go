package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"tiktok/config"
	"tiktok/internal/controller"
)

func StartServer() {
	gin.SetMode(config.Conf.App.Mode)
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	api := router.Group("/api")
	api.POST("/login", controller.Login)
	api.POST("/register", controller.Register)

	log.Fatal(router.Run(fmt.Sprintf(":%s", config.Conf.App.Port)))
}
