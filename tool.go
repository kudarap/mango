package puto

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

// ParsePayload parse post json input
func ParsePayload(r *http.Request, model interface{}) error {
	d := json.NewDecoder(r.Body)
	err := d.Decode(&model)

	return err
}

// ParseOption parse get params
func ParseOption(r *http.Request, model interface{}) error {
	d := schema.NewDecoder()
	err := d.Decode(model, r.URL.Query())

	return err
}

// GetSegment extract Request URI
func GetSegment(r *http.Request, key string) string {
	vars := mux.Vars(r)

	return vars[key]
}

// GetID on segment
func GetID(r *http.Request) (int, error) {
	id := GetSegment(r, "id")
	if id == "" {
		return 0, errors.New("id not exists")
	}

	// convert id to int
	i, err := strconv.Atoi(id)
	if err != nil {
		return 0, err
	}

	return i, nil
}
