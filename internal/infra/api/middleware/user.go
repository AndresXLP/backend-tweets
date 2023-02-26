package middleware

import (
	"context"
	"net/http"

	"github.com/andresxlp/backend-twitter/internal/domain/ports/postgres/repo"
	"github.com/andresxlp/backend-twitter/internal/utils"
	"github.com/labstack/echo/v4"
)

type UserMiddleware interface {
	OnlyUsers(next echo.HandlerFunc) echo.HandlerFunc
}

type onlyMiddleware struct {
	repo repo.Repository
	jwt  utils.JWT
}

func NewUserMiddleware(repo repo.Repository, jwt utils.JWT) UserMiddleware {
	return &onlyMiddleware{
		repo,
		jwt,
	}
}

func (u *onlyMiddleware) OnlyUsers(next echo.HandlerFunc) echo.HandlerFunc {
	return func(cntx echo.Context) error {
		ctx := cntx.Request().Context()

		token := cntx.Request().Header.Get("Authorization")
		if len(token) < 7 || token == "" || token[:6] != "Bearer" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Token not valid")
		}

		email, err := u.jwt.ValidateToken(token[7:])
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		}

		entityUser, err := u.repo.GetUser(ctx, email)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		if entityUser.ID == 0 {
			return echo.NewHTTPError(http.StatusUnauthorized, "User Unauthorized")
		}

		ctx = context.WithValue(ctx, "userID", entityUser.ID)
		cntx.SetRequest(cntx.Request().WithContext(ctx))
		return next(cntx)
	}
}
