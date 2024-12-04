package api

import (
	"github.com/gin-gonic/gin"
	"tiktok/global"
	"tiktok/internal/controller"
	"tiktok/internal/middleware"
	"tiktok/internal/repo"
	"tiktok/internal/repo/cache"
	"tiktok/internal/repo/dao"
	"tiktok/internal/service"
)

func RegisterUserRouter(router *gin.RouterGroup) {
	publicRouter := router.Group("/user")
	privateRouter := router.Group("/user")
	privateRouter.Use(middleware.Auth())

	db := global.Db
	rdb := global.Redis
	svc := service.NewUserService(
		repo.NewUserRepo(dao.NewUserDao(db)),
		repo.NewFollowRepo(dao.NewFollowDao(db), cache.NewFollowCache(rdb)),
	)
	ctl := controller.NewUserController(svc)
	{
		publicRouter.POST("/login", ctl.Login)
		publicRouter.POST("/register", ctl.Register)
		privateRouter.GET("/info", ctl.GetUserInfo)
	}
}
