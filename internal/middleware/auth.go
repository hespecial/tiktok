package middleware

import (
	"github.com/gin-gonic/gin"
	"tiktok/common/enum"
	"tiktok/common/response"
	"tiktok/pkg/jwt"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader(enum.JwtTokenName)
		if token == "" {
			response.Unauthorized(c)
			c.Abort()
			return
		}

		claims, err := jwt.ParseToken(token)
		if err != nil {
			response.Unauthorized(c)
			c.Abort()
			return
		}

		c.Set(enum.ContextUserId, claims.UserId)
		c.Next()
	}
}
