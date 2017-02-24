package store

import (
	"golang.org/x/net/context"
)

// Store data access
type Store interface {
	GetFoo(string) (string, error)
}

// GetFoo data
func GetFoo(c context.Context, id string) (string, error) {
	// return FromContext(c).GetFoo(id)
	return FromContext(c).GetFoo(id)
}
