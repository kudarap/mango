package test

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/javinc/puto/test/resource"
	"github.com/javinc/puto/test/service"
)

var (
	// Model for import use
	Model   = resource.Model{}
	decoder = schema.NewDecoder()
)

// Handler catches /test endpoint
func Handler(w http.ResponseWriter, r *http.Request) {
	log.Println("controllers TestHandler")

	// extract field from get params
	o := new(resource.Options)
	err := decoder.Decode(o, r.URL.Query())
	if err != nil {
		log.Panic(err)
	}

	log.Println("options", o)

	// extract field from post data
	err = r.ParseForm()
	if err != nil {
		log.Panic(err)
	}

	p := new(resource.Model)
	err = decoder.Decode(p, r.PostForm)
	if err != nil {
		log.Panic(err)
	}

	log.Println("payload", p)

	// lets base on the request type
	// use service
	var b []byte
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

	w.Write(b)
}
