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

	// extract field from get params
	o := new(resource.Options)
	err := x.Decoder.Decode(o, r.URL.Query())
	if err != nil {
		log.Panic(err)
	}

	log.Println("options", o)

	// parse json post
	decoder := json.NewDecoder(r.Body)
	p := new(resource.Model)
	decoder.Decode(&p)

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

func parseJSON(rw http.ResponseWriter, req *http.Request) {

}
