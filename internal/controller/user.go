package controller

import (
	"github.com/gin-gonic/gin"
	"tiktok/common/response"
	"tiktok/internal/repo/dao"
	"tiktok/internal/service"
	"tiktok/pkg/jwt"
	"tiktok/util"
)

func Login(c *gin.Context) {
	var form struct {
		Username string `form:"username" binding:"required" json:"username"`
		Password string `form:"password" binding:"required" json:"password"`
	}
	if err := c.ShouldBind(&form); err != nil {
		response.InvalidParams(c, err.Error())
		return
	}

	userService := service.NewUserService(dao.NewUserDao(dao.Db))

	user, err := userService.GetUserByUsername(form.Username)
	if err != nil {
		response.BadRequest(c)
		return
	}
	if user.Id == 0 {
		response.InvalidRequest(c, "username or password error")
		return
	}
	if util.Encrypt(form.Password) != user.Password {
		response.InvalidRequest(c, "username or password error")
		return
	}

	token, err := jwt.GenerateToken(user.Id)
	if err != nil {
		response.BadRequest(c)
		return
	}

	data := map[string]interface{}{
		"token": token,
	}
	response.Success(c, data)
}

func Register(c *gin.Context) {
	var form struct {
		Username string `form:"username" binding:"required" json:"username"`
		Password string `form:"password" binding:"required" json:"password"`
	}
	if err := c.ShouldBind(&form); err != nil {
		response.InvalidParams(c, err.Error())
		return
	}

	userService := service.NewUserService(dao.NewUserDao(dao.Db))

	user, err := userService.GetUserByUsername(form.Username)
	if err != nil {
		response.BadRequest(c)
		return
	}
	if user.Id != 0 {
		response.InvalidRequest(c, "username exists")
		return
	}

	if err = userService.CreateUser(form.Username, form.Password); err != nil {
		response.BadRequest(c)
		return
	}

	response.Success(c, nil)
}
