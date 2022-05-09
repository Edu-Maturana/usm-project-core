package utils

import (
	"back-usm/internals/auth/core/domain"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret string = GetEnvVar("JWT_SECRET")

func GenerateJWT(user domain.Admin) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 365).Unix(),
	})
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
