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

func PasswordCompare(p1, p2 string) error {
	return bcrypt.CompareHashAndPassword([]byte(p1), []byte(p2))
}
