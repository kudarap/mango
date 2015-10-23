package puto

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/schema"
)

// ParsePayload parse post json input
func ParsePayload(model interface{}, r *http.Request) error {
	d := json.NewDecoder(r.Body)
	err := d.Decode(&model)

	return err
}

// ParseOption parse get params
func ParseOption(model interface{}, r *http.Request) error {
	d := schema.NewDecoder()
	err := d.Decode(model, r.URL.Query())

	return err
}
