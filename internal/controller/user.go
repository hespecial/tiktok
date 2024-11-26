package controller

import (
	"github.com/gin-gonic/gin"
	"tiktok/common/code"
	"tiktok/common/response"
)

func Login(c *gin.Context) {
	var form struct {
		Username string `form:"username" binding:"required" json:"username"`
		Password string `form:"password" binding:"required" json:"password"`
	}
	if err := c.ShouldBind(&form); err != nil {
		response.Handle(code.InvalidParams, nil)
		return
	}
}
