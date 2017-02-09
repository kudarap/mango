package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/javinc/mango/tool"
)

// Middleware module
func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// default header
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.Header("Access-Control-Allow-Methods", "GET,POST,PATCH,DELETE,OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type,Authorization")

		// restrict allowed client when live
		// if Config.Live {
		// c.Header("Access-Control-Allow-Origin", Config.AppHost)
		// } else {
		c.Header("Access-Control-Allow-Origin", "*")
		// }

		// NOTE handle OPTIONS and HEAD method to respond immedietly
		if _, ok := tool.InArray(c.Request.Method, []string{
			http.MethodOptions,
			http.MethodHead,
		}); ok {
			c.AbortWithStatus(http.StatusOK)
		}

		c.Next()
	}
}
