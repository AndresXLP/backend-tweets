package dto

import "github.com/go-playground/validator/v10"

var validate = validator.New()

type User struct {
	Name     string `json:"name" validate:"required"`
	LastName string `json:"last_name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Address  string `json:"address" validate:"required"`
	Gender   string `json:"gender" validate:"required"`
	Age      int    `json:"age" validate:"required"`
}

type NewUser struct {
	User
	Password string `json:"password" validate:"required"`
}

func (u *NewUser) Validate() error {
	return validate.Struct(u)
}
