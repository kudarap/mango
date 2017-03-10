package middleware

import (
	"github.com/gin-gonic/gin"
)

// Middleware base
func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json; charset=utf-8")
	}
}
