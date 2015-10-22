package index

import (
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/gorilla/schema"
)

var decoder = schema.NewDecoder()

// Handler just to catch / uri
func Handler(w http.ResponseWriter, r *http.Request) {
	render := r.Method + " Goryo!\n"
	render += "URL is " + r.URL.String() + "\n"

	getParam := r.URL.Query()
	// get GET params
	render += "\nGET params: \n"
	for key, value := range getParam {
		render += "\t" + key + ":" + value[0] + "\n"
	}

	render += "\nPOST params: \n"
	// get POST params
	body, _ := ioutil.ReadAll(r.Body)
	postParam, _ := url.ParseQuery(string(body))
	for key, value := range postParam {
		render += "\t" + key + ":" + value[0] + "\n"
	}

	w.Write([]byte(render))
}
