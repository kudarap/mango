package test

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/javinc/mango/module"
)

// Login test
type Login struct {
	User string `json:"user" binding:"required"`
	Pass string `json:"pass" binding:"required"`
}

var service Service

// Handler test
func Handler(c *gin.Context) {
	module.SetContext(c)

	id := c.Param("id")

	switch c.Request.Method {
	case http.MethodGet:
		// detail
		if id != "" {
			d := service.Get(id)
			module.Output(d)

			return
		}

		// list
		d := service.Find()
		module.Output(d)

		return
	case http.MethodPost:
		var payload Object
		err := c.BindJSON(&payload)
		if err != nil {
			module.Panic(
				"REQUIRED_FIELDS",
				"user and pass field is required",
			)

			return
		}

		d := service.Create(payload)

		module.Output(d)

		return
	case http.MethodPut:
		if id == "" {
			module.Error(
				"RESOURCE_ID_REQUIRED",
				"resource id is missing",
			)

			return
		}

		module.Output(gin.H{
			"msg": "welcome PUT",
		})

		return
	case http.MethodDelete:
		if id == "" {
			module.Error(
				"RESOURCE_ID_REQUIRED",
				"resource id is missing",
			)

			return
		}

		module.Output(gin.H{
			"msg": "welcome DELETE",
		})

		return
	}

	module.Error(
		"METHOD_NOT_ALLOWED",
		c.Request.Method+" method not allowed in this endpoint",
	)
}
