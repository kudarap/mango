package middleware

import (
	"github.com/gin-gonic/gin"
)

// Metrics middleware accumulate request metrics
func Metrics() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}
