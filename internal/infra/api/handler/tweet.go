package handler

import (
	"net/http"

	"github.com/andresxlp/backend-twitter/internal/app"
	"github.com/andresxlp/backend-twitter/internal/domain/dto"
	"github.com/andresxlp/backend-twitter/internal/domain/entity"
	"github.com/labstack/echo/v4"
)

type Tweets interface {
	CreateTweet(cntx echo.Context) error
	GetTweets(cntx echo.Context) error
}

type tweets struct {
	app app.Tweets
}

func NewTweetsHandler(app app.Tweets) Tweets {
	return &tweets{app}
}

func (handler *tweets) CreateTweet(cntx echo.Context) error {
	ctx := cntx.Request().Context()

	tweet := dto.Tweet{}
	if err := cntx.Bind(&tweet); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := tweet.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := handler.app.CreateTweet(ctx, tweet); err != nil {
		return err
	}

	return cntx.JSON(http.StatusCreated, entity.Message{
		Message: "Tweet Created Successfully",
		Data:    tweet.Content,
	})
}

func (handler *tweets) GetTweets(cntx echo.Context) error {
	ctx := cntx.Request().Context()

	request := dto.TweetsRequest{}
	if err := cntx.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := request.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	pagi, err := handler.app.GetTweets(ctx, request)
	if err != nil {
		return err
	}

	return cntx.JSON(http.StatusOK, entity.Message{
		Message: "Tweets load successfully",
		Data:    pagi,
	})
}
