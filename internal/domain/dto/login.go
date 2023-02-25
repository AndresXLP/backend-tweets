package dto

type Login struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (l *Login) Validate() error {
	return validate.Struct(l)
}