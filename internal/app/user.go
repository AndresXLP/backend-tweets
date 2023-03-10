package app

import (
	"context"
	"net/http"

	"github.com/andresxlp/backend-twitter/internal/domain/dto"
	"github.com/andresxlp/backend-twitter/internal/domain/ports/postgres/repo"
	"github.com/andresxlp/backend-twitter/internal/infra/adapters/postgres/models"
	"github.com/labstack/echo/v4"
)

type User interface {
	CreateUser(ctx context.Context, newUser dto.NewUser) error
}

type user struct {
	userRepo repo.Repository
}

func NewUserApp(userRepo repo.Repository) User {
	return &user{userRepo}
}

func (app *user) CreateUser(ctx context.Context, newUser dto.NewUser) error {
	entityUser, err := app.userRepo.GetUser(ctx, newUser.Email)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if entityUser.ID != 0 {
		return echo.NewHTTPError(http.StatusConflict, "this email already register")
	}

	var userModel models.User
	userModel.BuildModel(newUser)
	if err = app.userRepo.CreateUser(ctx, userModel); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return nil
}
