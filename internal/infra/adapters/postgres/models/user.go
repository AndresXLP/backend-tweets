package models

import (
	"github.com/andresxlp/backend-twitter/internal/domain/dto"
)

type User struct {
	ID       int
	Name     string
	LastName string
	Email    string
	Address  string
	Gender   string
	Age      int
	Password string
}

func (u *User) BuildModel(newUser dto.NewUser) {
	u.Name = newUser.Name
	u.LastName = newUser.LastName
	u.Email = newUser.Email
	u.Address = newUser.Address
	u.Gender = newUser.Gender
	u.Age = newUser.Age
	u.Password = newUser.Password
}

func (u User) TableName() string {
	return "users"
}
