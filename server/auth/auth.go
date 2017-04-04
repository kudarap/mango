package auth

import (
	"errors"

	jwt "github.com/dgrijalva/jwt-go"
)

const salt = "andpepper"

// CheckToken checks the validity of the token and returns the payload
func CheckToken(token string) (map[string]interface{}, error) {
	if token == "" {
		return nil, errors.New("authorization token required")
	}

	c, err := checkToken(token)
	if err != nil {
		return nil, errors.New("invalid authorization token")
	}

	return c, nil
}

// CreateToken generates sign token
func CreateToken(payload map[string]interface{}) (string, error) {
	c := jwt.MapClaims{}
	// this is legal since 1.8 same struct signature
	c = payload
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	ts, err := token.SignedString([]byte(salt))

	return ts, err
}

func checkToken(s string) (jwt.MapClaims, error) {
	t, err := jwt.Parse(s, func(t *jwt.Token) (interface{}, error) {
		return []byte(salt), nil
	})

	if err != nil {
		return jwt.MapClaims{}, err
	}

	claims, _ := t.Claims.(jwt.MapClaims)

	return claims, err
}
