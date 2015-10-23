package puto

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/schema"
)

// ParsePayload parse post json input
func ParsePayload(model interface{}, r *http.Request) error {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&model)

	return err
}

// ParseOption parse get params
func ParseOption(model interface{}, r *http.Request) error {
	decoder := schema.NewDecoder()
	err := decoder.Decode(&model, r.URL.Query())

	return err
}
