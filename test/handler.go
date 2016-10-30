package test

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	x "github.com/javinc/mango/module"
)

var service Service

// Handler test
func Handler(c *gin.Context) {
	id := c.Param("id")
	x.SetContext(c)

	switch c.Request.Method {
	case http.MethodGet:
		// list
		if id == "" {
			filter := Object{
				Title:       c.Query("filter.title"),
				Description: c.Query("filter.description"),
			}

			option := Option{
				Slice:  c.Query("slice"),
				Order:  c.Query("order"),
				Filter: filter,
			}

			d, err := service.Find(option)
			if err != nil {
				x.Error("GET_RESOURCE_"+strings.ToUpper(resourceName), err.Error())
			}

			x.Output(d)

			return
		}

		// detail
		d, err := service.Get(id)
		if err != nil {
			x.Error("GET_RESOURCE_"+strings.ToUpper(resourceName), err.Error())

			return
		}

		x.Output(d)

		return
	case http.MethodPost:
		var payload Object
		err := c.BindJSON(&payload)
		if err != nil {
			x.Panic("REQUIRED_FIELDS", "field is required")

			return
		}

		d, err := service.Create(payload)
		if err != nil {
			x.Error("POST_RESOURCE_"+strings.ToUpper(resourceName), err.Error())
		}

		x.Output(d)

		return
	case http.MethodPatch:
		if id == "" {
			x.Error("RESOURCE_ID_REQUIRED", "resource id is missing")

			return
		}

		var payload Object
		err := c.BindJSON(&payload)
		if err != nil {
			x.Panic("REQUIRED_FIELDS", "field is required")

			return
		}

		d, err := service.Update(payload, id)
		if err != nil {
			x.Error("PUT_RESOURCE_"+strings.ToUpper(resourceName), err.Error())

			return
		}

		x.Output(d)

		return
	case http.MethodDelete:
		if id == "" {
			x.Error("RESOURCE_ID_REQUIRED", "resource id is missing")

			return
		}

		d, err := service.Remove(id)
		if err != nil {
			x.Error("DELETE_RESOURCE_"+strings.ToUpper(resourceName), err.Error())

			return
		}

		x.Output(d)

		return
	}

	x.MethodNotAllowed()
}
