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

