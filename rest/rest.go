package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	// GET method
	GET = http.MethodGet
	// POST method
	POST = http.MethodPost
	// PATCH method
	PATCH = http.MethodPatch
	// DELETE method
	DELETE = http.MethodDelete
)

// Router gin
func Router() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	r.Use(gin.Logger())

	return r
}
