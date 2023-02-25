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
	sessionRepo repo.Repository
	jwt         utils.JWT
	bcrypt      utils.Bcrypt
}

func NewSessionApp(sessionRepo repo.Repository, jwt utils.JWT, bcrypt utils.Bcrypt) Session {
	return &session{
		sessionRepo,
		jwt,
		bcrypt,
	}
}

func (app *session) Login(ctx context.Context, loginData dto.Login) (string, error) {
	entityUser, err := app.sessionRepo.GetUser(ctx, loginData.Email)
	if err != nil {
		return "", echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if entityUser.ID == 0 {
		return "", echo.NewHTTPError(http.StatusNotFound, "Wrong Email or Password")
	}

	if !(app.bcrypt.ValidatePassword(entityUser.Password, loginData.Password)) {
		return "", echo.NewHTTPError(http.StatusNotFound, "Wrong Email or Password")
	}

	token, err := app.jwt.GenerateToken(entityUser.Email)
	if err != nil {
		return "", echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return token, nil
}
