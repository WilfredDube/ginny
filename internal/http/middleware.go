package http

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func SecureHeader() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("X-XSS-Protection", "1; mode=block")
		c.Header("X-Frame-Options", "deny")

		c.Next()
	}
}

func RecoverPanic() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				ctx.Header("Connection", "close")
				ServerError(ctx, fmt.Errorf("%s", err))
			}
		}()

		ctx.Next()
	}
}
