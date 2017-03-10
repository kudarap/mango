package middleware

import (
	"github.com/gin-gonic/gin"
)

// DebugMiddleware accumulate request metrics
func DebugMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}
