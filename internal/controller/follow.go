package controller

import (
	"github.com/gin-gonic/gin"
	"tiktok/common/response"
	"tiktok/internal/repo"
	"tiktok/internal/service"
)

type FollowController struct {
	service service.FollowService
}

func NewFollowController(service service.FollowService) *FollowController {
	return &FollowController{
		service: service,
	}
}

func (fc *FollowController) DealRelation(c *gin.Context) {
	var form struct {
		TargetId int64 `form:"target_id" json:"target_id" binding:"required,min=1"`
		Action   uint8 `form:"action" json:"action" binding:"required,oneof=1 2"` // 1.follow 2.unfollow
	}
	if err := c.ShouldBind(&form); err != nil {
		response.InvalidParams(c, err)
		return
	}

	followService := service.NewFollowService(repo.Db)

	var err error
	userId := getUserId(c)
	if form.Action == 1 {
		err = followService.Follow(userId, form.TargetId)
	} else {
		err = followService.Unfollow(userId, form.TargetId)
	}
	if err != nil {
		response.BadRequest(c, err)
		return
	}

	response.Success(c, nil)
}

func (fc *FollowController) GetFollowingList(c *gin.Context) {
	userId := getUserId(c)
	followService := service.NewFollowService(repo.Db)

	list, err := followService.GetFollowingList(userId)
	if err != nil {
		response.BadRequest(c, err)
		return
	}

	response.Success(c, list)
}

func (fc *FollowController) GetFollowerList(c *gin.Context) {
	userId := getUserId(c)
	followService := service.NewFollowService(repo.Db)

	list, err := followService.GetFollowerList(userId)
	if err != nil {
		response.BadRequest(c, err)
		return
	}

	response.Success(c, list)
}
