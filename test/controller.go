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

	var err error
	var single bool
	var model interface{}

	// inits objects
	o := new(resource.Options)
	p := new(resource.Model)

	// extracting resource id
	if id, e := x.GetID(r); e == nil {
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
	switch r.Method {
	case "GET":
		// check if singles
		if single {
			model, err = service.Get(o)
		} else {
			model, err = service.Find(o)
		}
	case "POST":
		model, err = service.Create(p)
	case "DELETE":
		model, err = service.Remove(o)
	case "PUT":
		model, err = service.Update(p, o)
	}

	// response error
	if err != nil {
		x.ErrorOutput(w, err, 400)

		return
	}

	x.Output(w, model)
}
