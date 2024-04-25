package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func GeneratePassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}

	return string(hashed), nil
}

func CheckHashPassword(password, pwreq string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(pwreq), []byte(password))
	return err == nil
}
