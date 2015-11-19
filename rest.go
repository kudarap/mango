package puto

import (
	"encoding/json"
	"log"
	"net/http"
)

// error structure
type restError struct {
	Msg   string `json:"msg"`
	Panic bool   `json:"panic"`
}

// Output results
func Output(w http.ResponseWriter, model interface{}) {
	b, err := json.Marshal(model)

	// render some value
	if err != nil {
		log.Panic(err)

		return
	}

	render(w, b, 200)
}

// ErrorOutput with coes
func ErrorOutput(w http.ResponseWriter, e error, code int) {
	b, _ := json.Marshal(restError{
		e.Error(),
		false,
	})

	if code == 0 {
		code = 400
	}

	render(w, b, code)
}

func setHeaders(w http.ResponseWriter, code int) {
	w.Header().Set("access-control-allow-origin", "*")
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(code)
}

func render(w http.ResponseWriter, b []byte, code int) {
	setHeaders(w, code)
	w.Write(b)
}
