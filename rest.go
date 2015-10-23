package puto

import (
	"net/http"
)

// Render output
func Render(buffer []byte, w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Write(buffer)
}
