package user

import (
	"errors"

	"github.com/javinc/mango/module"
)

// Service object
type Service struct {
}

var resource Resource

// Find user
func (s *Service) Find(o Option) ([]Object, error) {
	return resource.Find(o)
}

// FindOne user
func (s *Service) FindOne(o Option) (Object, error) {
	return resource.FindOne(o)
}

// Get user
func (s *Service) Get(id string) (Object, error) {
	return resource.Get(id)
}

// Create user
func (s *Service) Create(p Object) (Object, error) {
	// check email existence
	user, _ := s.FindOne(Option{
		Filter: Object{
			Email: p.Email,
		},
	})

	if user.ID != "" {
		return Object{}, errors.New("email already exists")
	}

	// pasword hashing
	p.Password = module.Hash(p.Password)

	return resource.Create(p)
}

// Update user
func (s *Service) Update(p Object, id string) (Object, error) {
	return resource.Update(p, id)
}

// Remove user
func (s *Service) Remove(id string) (Object, error) {
	return resource.Remove(id)
}
