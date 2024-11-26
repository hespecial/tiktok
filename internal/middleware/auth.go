package middleware

import (
	"github.com/gin-gonic/gin"
	"tiktok/common/code"
	"tiktok/common/response"
	"tiktok/config"
	"tiktok/pkg/jwt"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader(config.JwtTokenName)
		if token == "" {
			response.Handle(code.Unauthorized, nil)
			c.Abort()
			return
		}

		claims, err := jwt.ParseToken(token)
		if err != nil {
			response.Handle(code.Unauthorized, nil)
			c.Abort()
			return
		}

		c.Set(config.ContextUserId, claims.UserId)
		c.Next()
	}
}
