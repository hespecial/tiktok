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

func RegisterFollowRouter(router *gin.RouterGroup) {
	//publicRouter := router.Group("/follow")
	privateRouter := router.Group("/relation")
	privateRouter.Use(middleware.Auth())

	db := global.Db
	rdb := global.Redis
	svc := service.NewFollowService(
		repo.NewFollowRepo(dao.NewFollowDao(db), cache.NewFollowCache(rdb)),
		repo.NewUserRepo(dao.NewUserDao(db)),
	)
	ctl := controller.NewFollowController(svc)
	{
		privateRouter.POST("/action", ctl.DealRelation)
		privateRouter.GET("/following", ctl.GetFollowingList)
		privateRouter.GET("/follower", ctl.GetFollowerList)
	}
}
