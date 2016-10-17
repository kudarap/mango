package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

// port server
const port = "8000"

func main() {
	router := gin.New()

	// Routes consist of a path and a handler function.
	router.Any("/", indexHandler)

	log.Println("serving on port", port)
	router.Run(":" + port)
}

func indexHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
