package response

import (
	"log/slog"
	"tiktok/common/code"
)

type Response struct {
	Code    code.Code   `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Handle(c code.Code, data interface{}, message ...string) *Response {
	msg := c.GetMessage()
	if len(msg) > 0 {
		msg = message[0]
	}

	if c != code.Success {
		slog.Info(msg)
	}

	return &Response{
		Code:    c,
		Message: msg,
		Data:    data,
	}
}
