package user

import x "github.com/javinc/mango/module"

// Service object
type Service struct {
}

// EmailAuth email authentication
type EmailAuth struct {
	ID    string `json:"id"`
	Type  string `json:"type"`
	Token string `json:"token"`
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
func (t *Service) Login(email, pass string) (EmailAuth, error) {
	o := Option{
		Filter: Object{
			Email: email,
		},
	}

	user, err := resource.FindOne(o)
	if err != nil {
		return EmailAuth{}, err
	}

	auth := EmailAuth{
		ID:   user.ID,
		Type: "admin",
	}

	token, err := x.CreateToken(map[string]interface{}{
		"id":   auth.ID,
		"type": auth.Type,
	})

	if err != nil {
		return EmailAuth{}, err
	}

	auth.Token = token

	return auth, nil
}
