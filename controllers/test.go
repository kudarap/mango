package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	res "github.com/javinc/puto/resources/test"
	"github.com/javinc/puto/services/test"
)

// TestHandler catches /test endpoint
func TestHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("controllers TestHandler")

	// extract field from get params
	o := new(res.Options)
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

	p := new(res.Model)
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
		b, err = json.Marshal(test.Find(o))

	case "POST":
		b, err = json.Marshal(test.Create(p))
	}

	// render some value
	if err != nil {
		log.Panic(err)
	}

	w.Write(b)
}
