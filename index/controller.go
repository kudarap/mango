package index

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// Handler just to catch / uri
func Handler(w http.ResponseWriter, r *http.Request) {
	log.Println("index handler")

	render := "Method: " + r.Method + "\n\n"
	render += "URL: " + r.URL.String() + "\n"

	getParam := r.URL.Query()
	// get GET params
	render += "\nGET: \n"
	for key, value := range getParam {
		render += "\t" + key + ":" + value[0] + "\n"
	}

	render += "\nPOST: \n"
	// get POST params
	body, _ := ioutil.ReadAll(r.Body)
	postParam, _ := url.ParseQuery(string(body))
	for key, value := range postParam {
		render += "\t" + key + ":" + value[0] + "\n"
	}

	w.Write([]byte(render))
}
