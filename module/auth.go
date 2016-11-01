package module

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

const appKey = "blinkdagger"

// CreateToken generates JWT
func CreateToken(payload map[string]interface{}) (string, error) {
	claims := jwt.MapClaims{}

	for k, v := range payload {
		claims[k] = v
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ts, err := token.SignedString([]byte(appKey))

	fmt.Println(ts)

	return ts, err
}

func check(ts string) {
	token, err := jwt.Parse(ts, func(token *jwt.Token) (interface{}, error) {
		return []byte(appKey), nil
	})

	claims, ok := token.Claims.(jwt.MapClaims)

	fmt.Println(claims)
	fmt.Println(ok)
	fmt.Println(err)
}

func blink() {

}
