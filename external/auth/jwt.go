package external

// this package mimic external library
// assume api key use jwt
// and algorithm used is HS256
// and have additional calim "role" (for unauthorized error CMIIW)

import (
	"fmt"

	"github.com/golang-jwt/jwt"
)

const (
	Admin = iota + 1
	User
)

var jwtSecret = "secret" // TODO: use config

type JwtBusinessClaim struct {
	Role int `json:"role"`
	jwt.StandardClaims
}

func ValidateJWT(tokenString string) (interface{}, error) {
	claims := JwtBusinessClaim{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		if t.Method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("external.ValidateJWT: miss match algo type got %s want 'HS256:'", t.Method.Alg())
		}

		return []byte(jwtSecret), nil
	})

	if err != nil {
		err = fmt.Errorf("external.ValidateJWT: error parse claims %w", err)
		return nil, err
	}

	return token, nil
}
