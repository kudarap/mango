package puto

import (
	"encoding/json"
	"net/http"
)

// error structure
type error struct {
	Message string
	Panic   bool
}

// Output results
func Output(w http.ResponseWriter, b []byte) {
	setHeaders(w)
	w.Write(b)
}

// ErrorOutput with coes
func ErrorOutput(w http.ResponseWriter, e error, code int) {
	setHeaders(w)

	b, _ := json.Marshal(error{
		e.Error(),
		false,
	})

	http.Error(w, string(b), code)
}

func setHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
}
