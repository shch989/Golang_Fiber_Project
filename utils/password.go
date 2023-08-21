package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashingPassword(password string) (string, error) {
	hashedByte, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(hashedByte), nil
}

func CheckPasswordHash(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		log.Println(err.Error())
		log.Println("password : ", password)
		log.Println("hashedPassword : ", hashedPassword)
	}
	return err == nil
}
