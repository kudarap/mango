package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/javinc/mango/server/auth"
)

// PrivateMiddleware checks if has valid token
func PrivateMiddleware(checkPayload func(map[string]interface{}) error) gin.HandlerFunc {
	return func(c *gin.Context) {
		// validates token
		p, err := auth.CheckToken(
			getToken(c.Request.Header.Get("authorization")))
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"name":    "AUTH_REQUIRED",
				"message": err.Error(),
			})
			c.Abort()

			return
		}

		// validates payload
		err = checkPayload(p)
		if err != nil {
			// NOTE you can check use 404
			c.JSON(http.StatusForbidden, gin.H{
				"name":    "AUTH_INVALID",
				"message": err.Error(),
			})
			c.Abort()

			return
		}

		c.Next()
	}
}

func getToken(authHeader string) string {
	s := strings.Split(strings.TrimSpace(authHeader), " ")
	if len(s) != 2 {
		return ""
	}

	return s[1]
}
