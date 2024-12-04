package controller

import (
	"github.com/gin-gonic/gin"
	"tiktok/common/enum"
	"tiktok/common/response"
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
		TargetId int64         `form:"target_id" json:"target_id" binding:"required,min=1"`
		Action   enum.Relation `form:"action" json:"action" binding:"oneof=0 1"` // 0.unfollow 1.follow
	}
	if err := c.ShouldBind(&form); err != nil {
		response.InvalidParams(c, err)
		return
	}

	var err error
	userId := getUserId(c)
	if form.Action == enum.RelationFollow {
		err = fc.service.Follow(c, userId, form.TargetId)
	} else {
		err = fc.service.Unfollow(c, userId, form.TargetId)
	}
	if err != nil {
		response.BadRequest(c, err)
		return
	}

	response.Success(c)
}

func (fc *FollowController) GetFollowingList(c *gin.Context) {
	userId := getUserId(c)

	list, err := fc.service.GetFollowingList(c, userId)
	if err != nil {
		response.BadRequest(c, err)
		return
	}

	response.Success(c, list)
}

func (fc *FollowController) GetFollowerList(c *gin.Context) {
	userId := getUserId(c)

	list, err := fc.service.GetFollowerList(c, userId)
	if err != nil {
		response.BadRequest(c, err)
		return
	}

	response.Success(c, list)
}
