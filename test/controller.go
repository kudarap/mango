package test

import (
	"encoding/json"
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
	x.ParseOption(o, r)
	log.Println("options", o)

	// extract field from json post
	x.ParsePayload(p, r)
	log.Println("payload", p)

	// lets base on the request type then use the service
	var b []byte
	var err error
	switch r.Method {
	case "GET":
		b, err = json.Marshal(service.Find(o))

	case "POST":
		b, err = json.Marshal(service.Create(p))
	}

	// render some value
	if err != nil {
		log.Panic(err)
	}

	x.Render(b, w)
}
