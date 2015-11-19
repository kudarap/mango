package test

import (
	"log"
	"net/http"

	x "github.com/javinc/puto"
	"github.com/javinc/puto/test/resource"
	"github.com/javinc/puto/test/service"
)

var (
	// Model for import use
	Model = resource.Model{}
)

// Handler catches /test endpoint
func Handler(w http.ResponseWriter, r *http.Request) {
	log.Println("test handler")

	o := new(resource.Options)
	p := new(resource.Model)

	// extract field from get params
	x.GetOption(r, o)
	log.Println("options", o)

	// extract field from json post
	x.GetPayload(r, p)
	log.Println("payload", p)

	// lets base on the request type then use the service
	var err error
	var model interface{}
	// var err error
	switch r.Method {
	case "GET":
		id, err := x.GetID(r)
		if err != nil {
			model, err = service.Find(o)
		} else {
			o.Filters.ID = id
			model, err = service.Get(o)
		}
	case "POST":
		model, err = service.Create(p)
	case "DELETE":
		model, err = service.Remove(o)
	case "PUT":
		model, err = service.Update(p, o)
	}

	if err != nil {
		x.ErrorOutput(w, err, 400)
	} else {
		x.Output(w, model)
	}
}
