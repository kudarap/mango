package server

import (
	"github.com/gin-gonic/gin"
)

// Engine gin engine
func Engine(mode string) *gin.Engine {
	gin.SetMode(mode)

	r := gin.New()
	r.Use(gin.Recovery())

	return r
}
