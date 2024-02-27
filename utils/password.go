package utils

import (
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		fmt.Println("there is error in hashing password...", err)
		return "", errors.New("error in hashing password")
	}

	return string(hashedPassword), nil
}

func VerifyPassword(hashedPassword string, userPassword string) error{
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword),[]byte(userPassword))
}
