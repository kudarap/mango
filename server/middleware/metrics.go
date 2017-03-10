package middleware

import (
	"github.com/gin-gonic/gin"
)

// MetricsMiddleware accumulate request metrics
func MetricsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}
