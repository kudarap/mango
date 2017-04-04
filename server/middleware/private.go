package middleware

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/javinc/mango/server/auth"
)

// PrivateMiddleware checks if has valid token
func PrivateMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := getToken(c.Request.Header.Get("authorization"))

		p, err := auth.CheckToken(t)
		if err != nil {
			c.String(400, err.Error())
			c.Abort()

			return
		}

		log.Println("CHECKING IN", err, p)
	}
}

func getToken(authHeader string) string {
	s := strings.Split(strings.TrimSpace(authHeader), " ")
	if len(s) != 2 {
		return ""
	}

	return s[1]
}
