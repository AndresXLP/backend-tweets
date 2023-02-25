package utils

import "golang.org/x/crypto/bcrypt"

type Bcrypt interface {
	HashPassword(password *string)
	ValidatePassword(dbPassword []byte, userPassword string) bool
}

type bcryptUtil struct{}

func NewBcryptUtil() Bcrypt {
	return &bcryptUtil{}
}

func (b *bcryptUtil) HashPassword(password *string) {
	hash, _ := bcrypt.GenerateFromPassword([]byte(*password), bcrypt.DefaultCost)

	*password = string(hash)
}

func (b *bcryptUtil) ValidatePassword(dbPassword []byte, userPassword string) bool {
	err := bcrypt.CompareHashAndPassword(dbPassword, []byte(userPassword))
	return err == nil
}
