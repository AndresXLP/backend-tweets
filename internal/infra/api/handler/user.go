package handler

import (
	"net/http"

	"github.com/andresxlp/backend-twitter/internal/app"
	"github.com/andresxlp/backend-twitter/internal/domain/dto"
	"github.com/andresxlp/backend-twitter/internal/domain/entity"
	"github.com/andresxlp/backend-twitter/internal/utils"
	"github.com/labstack/echo/v4"
)

type User interface {
	CreateUser(cntx echo.Context) error
}

type user struct {
	app    app.User
	bcrypt utils.Bcrypt
}

func NewUserHandler(app app.User, bcrypt utils.Bcrypt) User {
	return &user{
		app,
		bcrypt,
	}
}

func (handler *user) CreateUser(cntx echo.Context) error {
	ctx := cntx.Request().Context()

	newUser := dto.NewUser{}
	if err := cntx.Bind(&newUser); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := newUser.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	handler.bcrypt.HashPassword(&newUser.Password)

	if err := handler.app.CreateUser(ctx, newUser); err != nil {
		return err
	}

	return cntx.JSON(http.StatusCreated, entity.Message{
		Message: "User Created Successfully",
		Data:    newUser.Email,
	})
}
