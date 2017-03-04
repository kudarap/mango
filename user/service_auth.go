package user

// Login user
import (
	"errors"

	x "github.com/javinc/mango/module"
)

// Login user via email
func (s *Service) Login(email, pass string) (map[string]interface{}, error) {
	payload := make(map[string]interface{})
	user, err := s.FindOne(Option{
		Filter: Object{
			Email:    email,
			Password: x.Hash(pass),
		},
	})

	if err != nil {
		return payload, errors.New("invalid email or password")
	}

	// payload population
	payload["id"] = user.ID
	payload["type"] = user.Type

	token, err := x.CreateToken(x.AuthUser{
		ID:   user.ID,
		Type: payload["type"].(string),
	})

	payload["token"] = token

	return payload, err
}

// Register user
func (s *Service) Register(p Object) (Object, error) {
	// testing instance
	if p.Email == "testlogin@mango.com" {
		p.ID = "testloginid"
	}

	return s.Create(p)
}
