package utils

import (
	"os"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(claims *jwt.StandardClaims) (string, error) {
	tokens := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwt, err := tokens.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return jwt, nil
}
