package file

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	x "github.com/javinc/mango/module"
)

var service Service

// Handler file
func Handler(c *gin.Context) {
	id := c.Param("id")
	switch c.Request.Method {
	case x.GET:
		// list
		if id == "" {
			size, _ := strconv.Atoi(c.Query("filter.size"))
			filter := Object{
				Slug:      c.Query("filter.slug"),
				Extension: c.Query("filter.extension"),
				Mime:      c.Query("filter.mime"),
				Size:      size,
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
	case x.POST:
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
	case x.PATCH:
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
	case x.DELETE:
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

	x.MethodNotAllowedError()
}
