package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CORS middleware injects CORS headers to each request
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Header("Access-Control-Allow-Methods", "HEAD, GET, POST, PATCH, DELETE, OPTIONS")

		// NOTE handle OPTIONS and HEAD method to respond immedietly
		if c.Request.Method == http.MethodHead || c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusOK)
		}

		c.Next()
	}
}
