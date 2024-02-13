package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func T() {
	fmt.Printf("kjjkf")
}
func HashPassword(password string) string {
	passwordCypted, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	return string(passwordCypted)
}
