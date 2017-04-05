package middleware

import (
	"github.com/gin-gonic/gin"
)

// Debug middleware accumulate request metrics
func Debug() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}
