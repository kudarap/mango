package module

import (
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

const appKey = "blinkdagger"

// AuthUser authenticated user data
type AuthUser struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

// CreateToken generates JWT
func CreateToken(payload AuthUser) (string, error) {
	claims := jwt.MapClaims{
		"id":   payload.ID,
		"type": payload.Type,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ts, err := token.SignedString([]byte(appKey))

	fmt.Println(ts)

	return ts, err
}

// GetAuth checking authentication with header
func GetAuth() (AuthUser, error) {
	authHeaders := strings.Split(strings.TrimSpace(ctx.Request.Header.Get("authorization")), " ")
	claims, err := checkToken(authHeaders[1])

	if err != nil {
		return AuthUser{}, err
	}

	return AuthUser{
		ID:   claims["id"].(string),
		Type: claims["type"].(string),
	}, err
}

func checkToken(ts string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(ts, func(token *jwt.Token) (interface{}, error) {
		return []byte(appKey), nil
	})

	if err != nil {
		return jwt.MapClaims{}, err
	}

	claims, _ := token.Claims.(jwt.MapClaims)

	return claims, err
}
