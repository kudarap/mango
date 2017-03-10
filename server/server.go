package server

import (
	"github.com/gin-gonic/gin"
)

// Engine gin engine
func Engine() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	r.Use(gin.Recovery())

	return r
}
