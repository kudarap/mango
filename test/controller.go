package test

import (
	"errors"
	"log"

	"github.com/gin-gonic/gin"
	x "github.com/javinc/puto"
	"github.com/javinc/puto/test/resource"
	"github.com/javinc/puto/test/service"
)

var (
	// Model for import use
	Model = resource.Model{}
)

// Handler catches /test endpoint
func Handler(c *gin.Context) {
	log.Println("test handler")

	var err error
	var single bool
	var model interface{}

	r := c.Request
	w := c.Writer

	// inits objects and load default values
	o := new(resource.Options)
	o.LoadDefaults()
	p := new(resource.Model)

	// extracting resource id
	if id, e := x.GetID(c); e == nil {
		// assign id on Filters
		single = true
		o.Filters.ID = id
	}

	// extract field from get params
	x.ParseOption(r, o)
	log.Println("id", o.Filters.ID)
	log.Println("options", o)

	// extract field from json post
	x.ParsePayload(r, p)
	log.Println("payload", p)

	// lets base on the request type then use the service
	err = errors.New("request method not allowed")
	switch r.Method {
	case "GET":
		if single {
			model, err = service.Get(o)
		} else {
			model, err = service.Find(o)
		}
	case "POST":
		model, err = service.Create(p)
	case "DELETE":
		err = errors.New("id not defined")
		if single {
			model, err = service.Remove(o)
		}
	case "PUT":
		err = errors.New("id not defined")
		if single {
			model, err = service.Update(p, o)
		}
	}

	// response error
	if err != nil {
		x.ErrorOutput(w, err, 400)

		return
	}

	x.Output(w, model)
}
