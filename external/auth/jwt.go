package external

// this package mimic external library
// assume api key use jwt
// and algorithm used is HS256

import (
	"fmt"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
)

const (
	Admin = iota + 1
	User
)

var jwtSecret = "secret"

type JwtBusinessClaim struct {
	Role int `json:"role"`
	jwt.StandardClaims
	testing bool
}

func (claim JwtBusinessClaim) IsTesting() bool {
	return claim.testing
}

// generate JwtBusinessClaim which bypass validateJWT()
func NewJwtTestingClaim() *JwtBusinessClaim {
	return &JwtBusinessClaim{testing: true}
}

func ValidateJWT(tokenString string) (interface{}, error) {
	var claims *JwtBusinessClaim
	if os.Getenv("BS_ENV") == "development" && os.Getenv("TEST_ENV") == "1" {
		claims = &JwtBusinessClaim{testing: true}
	} else {
		claims = &JwtBusinessClaim{}
	}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		if claims.IsTesting() {
			return []byte(jwtSecret), nil
		}
		if t.Method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("external.ValidateJWT: miss match algo type got %s want 'HS256:'", t.Method.Alg())
		}
		if strings.ToLower(claims.Issuer) != "62teknologi" {
			return nil, fmt.Errorf("external.ValidateJWT: unexpected issuer")
		}

		return []byte(jwtSecret), nil
	})

	if err != nil {
		err = fmt.Errorf("external.ValidateJWT: error parse claims %w", err)
		return nil, err
	}
	payload, ok := token.Claims.(*JwtBusinessClaim)
	if !ok {
		err = fmt.Errorf("service.businessService.Delete: invalid claim type %T", token)
		return nil, err
	}

	return payload, nil
}
