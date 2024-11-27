package response

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"tiktok/common/code"
)

type Response struct {
	Code    code.Code   `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Handle(ctx *gin.Context, c code.Code, data interface{}, message ...string) {
	msg := c.GetMessage()
	if len(message) > 0 {
		msg = message[0]
	}

	if c != code.Success {
		slog.Info(msg)
	}

	ctx.JSON(http.StatusOK, Response{
		Code:    c,
		Message: msg,
		Data:    data,
	})
}

func InvalidParams(ctx *gin.Context, message ...string) {
	Handle(ctx, code.InvalidParams, nil, message...)
}

func InvalidRequest(ctx *gin.Context, message ...string) {
	Handle(ctx, code.InvalidRequest, nil, message...)
}

func BadRequest(ctx *gin.Context, message ...string) {
	Handle(ctx, code.BadRequest, nil, message...)
}

func Unauthorized(ctx *gin.Context, message ...string) {
	Handle(ctx, code.Unauthorized, nil, message...)
}

func Success(ctx *gin.Context, data interface{}) {
	Handle(ctx, code.Success, data)
}
