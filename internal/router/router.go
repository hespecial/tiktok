package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"tiktok/config"
	"tiktok/internal/controller"
	"tiktok/internal/middleware"
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
	api.POST("/user/login", controller.Login)
	api.POST("/user/register", controller.Register)
	api.GET("/user/info", middleware.Auth(), controller.GetUserInfo)

	api.POST("/relation/action", middleware.Auth(), controller.DealRelation)
	api.GET("/relation/following", middleware.Auth(), controller.GetFollowingList)
	api.GET("/relation/follower", middleware.Auth(), controller.GetFollowerList)

	log.Fatal(router.Run(fmt.Sprintf(":%s", config.Conf.App.Port)))
}
