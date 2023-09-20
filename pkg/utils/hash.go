package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func PasswordHash(password string) string {
	if len(password) != 0 {
		bytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		return string(bytes)
	}
	return password
}

func PasswordCompare(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
