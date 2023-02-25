package app

import (
	"context"
	"net/http"

	"github.com/andresxlp/backend-twitter/internal/domain/dto"
	"github.com/andresxlp/backend-twitter/internal/domain/ports/postgres/repo"
	"github.com/andresxlp/backend-twitter/internal/utils"
	"github.com/labstack/echo/v4"
)

type Session interface {
	Login(ctx context.Context, loginData dto.Login) (string, error)
}

type session struct {
	sessionRepo repo.User
}

func NewSessionApp(sessionRepo repo.User) Session {
	return &session{sessionRepo}
}

func (app *session) Login(ctx context.Context, loginData dto.Login) (string, error) {
	entityUser, err := app.sessionRepo.GetUser(ctx, loginData.Email)
	if err != nil {
		return "", echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if entityUser.ID == 0 {
		return "", echo.NewHTTPError(http.StatusNotFound, "Wrong Email or Password")
	}

	if !(utils.ValidatePassword(entityUser.Password, loginData.Password)) {
		return "", echo.NewHTTPError(http.StatusNotFound, "Wrong Email or Password")
	}

	token, err := utils.GenerateToken(entityUser.Email)
	if err != nil {
		return "", echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return token, nil
}
