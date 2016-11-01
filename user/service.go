package user

import (
	"errors"

	x "github.com/javinc/mango/module"
)

// Service object
type Service struct {
}

var resource Resource

// Find test
func (t *Service) Find(o Option) ([]Object, error) {
	return resource.Find(o)
}

// Get test
func (t *Service) Get(id string) (Object, error) {
	return resource.Get(id)
}

// Create test
func (t *Service) Create(p Object) (Object, error) {
	return resource.Create(p)
}

// Update test
func (t *Service) Update(p Object, id string) (Object, error) {
	return resource.Update(p, id)
}

// Remove test
func (t *Service) Remove(id string) (Object, error) {
	return resource.Remove(id)
}

// Login user
func (t *Service) Login(email, pass string) (map[string]interface{}, error) {
	o := Option{
		Filter: Object{
			Email:    email,
			Password: x.Hash(pass),
		},
	}

	payload := make(map[string]interface{})

	user, err := resource.FindOne(o)
	if err != nil {
		return payload, errors.New("invalid email or password")
	}

	// payload population
	payload["id"] = user.ID
	payload["type"] = "admin"

	token, err := x.CreateToken(x.AuthUser{
		ID:   user.ID,
		Type: payload["type"].(string),
	})

	payload["token"] = token

	return payload, err
}
