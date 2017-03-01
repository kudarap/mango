package logic

import (
	"github.com/javinc/mango/model"
	"github.com/javinc/mango/service"
)

// datastore is an implementation of a model.Store built on top
// of the rethink database driver
type logic struct {
	User model.User
}

// New creates database connection
func New(u model.User) service.Service {
	return &logic{
		User: u,
	}
}
