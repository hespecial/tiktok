package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"tiktok/config"
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
	api.POST("/login")

	log.Fatal(router.Run(fmt.Sprintf(":%s", config.Conf.App.Port)))
}
