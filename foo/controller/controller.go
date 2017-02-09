package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/javinc/mango/foo/schema"
	"github.com/javinc/mango/foo/service"
)

// Handler http call
func Handler(c *gin.Context) {
	o := schema.Option{}

	service.Find(o)

	c.JSON(200, "hello world")
}
