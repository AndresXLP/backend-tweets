package app

import (
	"context"

	"github.com/andresxlp/backend-twitter/internal/domain/dto"
	"github.com/andresxlp/backend-twitter/internal/domain/ports/postgres/repo"
	"github.com/andresxlp/backend-twitter/internal/infra/adapters/postgres/models"
)

type User interface {
	CreateUser(ctx context.Context, newUser dto.NewUser) error
}

type user struct {
	userRepo repo.User
}

func NewUserApp(userRepo repo.User) User {
	return &user{userRepo}
}

func (app *user) CreateUser(ctx context.Context, newUser dto.NewUser) error {
	var userModel models.User
	userModel.BuildModel(newUser)
	if err := app.userRepo.CreateUser(ctx, userModel); err != nil {
		return err
	}
	return nil
}
