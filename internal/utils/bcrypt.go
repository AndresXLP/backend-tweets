package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) []byte {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return hash
}

func ValidatePassword(dbPassword []byte, userPassword string) bool {
	err := bcrypt.CompareHashAndPassword(dbPassword, []byte(userPassword))
	return err == nil
}
