package puto

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
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

// GetID on segment
func GetID(c *gin.Context) (int, error) {
	id := strings.Trim(c.Param("id"), "/")
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
