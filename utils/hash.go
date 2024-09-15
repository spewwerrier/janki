package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func Hash(value string) string {
	v, _ := bcrypt.GenerateFromPassword([]byte(value), 10)
	return string(v)
}

func CheckHash(hashed_password string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed_password), []byte(password))
	return err == nil
}
