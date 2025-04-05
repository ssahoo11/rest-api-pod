package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func GetHashPassword(passwd string) (string, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(passwd), 14)
	return string(password), err
}

func ComparePassword(password string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
