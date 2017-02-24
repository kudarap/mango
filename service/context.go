package service

import (
	"golang.org/x/net/context"
)

const key = "service"

// Setter defines a context that enables setting values.
type Setter interface {
	Set(string, interface{})
}

// FromContext returns the Service associated with this context.
func FromContext(c context.Context) Service {
	return c.Value(key).(Service)
}

// ToContext adds the Service to this context if it supports
// the Setter interface.
func ToContext(c Setter, store Service) {
	c.Set(key, store)
}
