package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	x "github.com/javinc/puto"
	"github.com/javinc/puto/index"
	"github.com/javinc/puto/test"
	"github.com/javinc/puto/user"
)

// port server
const port = "8000"

func sample(c *gin.Context) {
	name := c.Param("id")
	c.String(http.StatusOK, "Hello %s", name)
}

func main() {
	router := gin.New()

	// Routes consist of a path and a handler function.
	router.Any("/", index.Handler)
	router.Any("/test/*id", test.Handler)
	router.Any("/user/*id", user.Handler)

	// Migrates Db
	x.MySQL.AutoMigrate(
		&test.Model,
		&user.Model,
	)

	log.Println("serving on port", port)
	router.Run(":" + port)
}
