package middleware

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	if err != nil {
		log.Println(err.Error())
	}
	return string(bytes)
}
