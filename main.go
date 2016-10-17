package main

import (
	"github.com/gin-gonic/gin"
	"github.com/javinc/mango/test"
)

func main() {
	router := gin.New()

	// Routes consist of a path and a handler function.
	router.Any("/test", test.Handler)

	router.Run(":8000")
}
