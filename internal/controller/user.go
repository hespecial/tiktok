package controller

import (
	"github.com/gin-gonic/gin"
	"tiktok/common/enum"
	"tiktok/common/response"
	"tiktok/internal/service"
	"tiktok/pkg/jwt"
	"tiktok/util"
)

type UserController struct {
	service service.UserService
}

func NewUserController(service service.UserService) *UserController {
	return &UserController{
		service: service,
	}
}

func getUserId(c *gin.Context) int64 {
	userId, _ := c.Get(enum.ContextUserId)
	return userId.(int64)
}

func (uc *UserController) Login(c *gin.Context) {
	var form struct {
		Username string `form:"username" json:"username" binding:"required,min=4,max=30,alphanum"`
		Password string `form:"password" json:"password" binding:"required,min=6,max=128"`
	}
	if err := c.ShouldBind(&form); err != nil {
		response.InvalidParams(c, err)
		return
	}

	user, err := uc.service.GetUserByUsername(c, form.Username)
	if err != nil {
		response.BadRequest(c, err)
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
		response.BadRequest(c, err)
		return
	}

	data := map[string]interface{}{
		"token": token,
	}
	response.Success(c, data)
}

func (uc *UserController) Register(c *gin.Context) {
	var form struct {
		Username string `form:"username" json:"username" binding:"required,required,min=4,max=30,alphanum"`
		Password string `form:"password" json:"password" binding:"required,min=6,max=128"`
	}
	if err := c.ShouldBind(&form); err != nil {
		response.InvalidParams(c, err)
		return
	}

	user, err := uc.service.GetUserByUsername(c, form.Username)
	if err != nil {
		response.BadRequest(c, err)
		return
	}
	if user.Id != 0 {
		response.InvalidRequest(c, "username exists")
		return
	}

	if err = uc.service.CreateUser(c, form.Username, form.Password); err != nil {
		response.BadRequest(c, err)
		return
	}

	response.Success(c)
}

func (uc *UserController) GetUserInfo(c *gin.Context) {
	userId := getUserId(c)

	user, err := uc.service.GetUserInfo(c, userId)
	if err != nil {
		response.BadRequest(c, err)
		return
	}

	response.Success(c, user)
}
