package resource

import (
	"github.com/javinc/mango/model"
	"github.com/javinc/mango/store"
)

// datastore is an implementation of a model.Store built on top
// of the rethink database driver
type resource struct {
	User model.User
}

// New creates database connection
func New(u model.User) store.Store {
	return &resource{
		User: u,
	}
}
