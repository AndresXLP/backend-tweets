package handler

import (
	"net/http"

	"github.com/andresxlp/backend-twitter/internal/app"
	"github.com/andresxlp/backend-twitter/internal/domain/dto"
	"github.com/andresxlp/backend-twitter/internal/domain/entity"
	"github.com/labstack/echo/v4"
)

type Session interface {
	Login(cntx echo.Context) error
}

type session struct {
	app app.Session
}

func NewSessionHandler(app app.Session) Session {
	return &session{app}
}

//	@Tags			Login
//	@Summary		Login User
//	@Description	Login
//	@Produce		json
//	@Param			request	body		dto.Login	true	"Request Body"
//	@Success		200		{object}	entity.Message
//	@Failure		400
//	@Failure		404
//	@Router			/session/login [post]
func (handler *session) Login(cntx echo.Context) error {
	ctx := cntx.Request().Context()
	loginData := dto.Login{}
	if err := cntx.Bind(&loginData); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := loginData.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	token, err := handler.app.Login(ctx, loginData)
	if err != nil {
		return err
	}

	return cntx.JSON(http.StatusOK, entity.Message{
		Message: "Session Login Successfully",
		Data:    token,
	})
}
