package models

import "github.com/andresxlp/backend-twitter/internal/domain/dto"

type Login struct {
	Email    string
	Password string
}

func (l *Login) BuildModel(loginData dto.Login) {
	l.Email = loginData.Email
}
