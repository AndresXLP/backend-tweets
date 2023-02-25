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
	UpdateTweet(cntx echo.Context) error
	DeleteTweet(cntx echo.Context) error
}

type tweets struct {
	app app.Tweets
}

func NewTweetsHandler(app app.Tweets) Tweets {
	return &tweets{app}
}

func (handler *tweets) CreateTweet(cntx echo.Context) error {
	ctx := cntx.Request().Context()

	tweet := dto.Tweets{}
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

	request.SetDefault()

	pagination, err := handler.app.GetTweets(ctx, request)
	if err != nil {
		return err
	}

	return cntx.JSON(http.StatusOK, entity.Message{
		Message: "Tweets load successfully",
		Data:    pagination,
	})
}

func (handler *tweets) UpdateTweet(cntx echo.Context) error {
	ctx := cntx.Request().Context()

	updateRequest := dto.Tweets{}
	if err := cntx.Bind(&updateRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := updateRequest.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := handler.app.UpdateTweet(ctx, updateRequest); err != nil {
		return err
	}

	return cntx.JSON(http.StatusOK, entity.MessageSuccess{
		Message: "Update tweet Successfully",
	})
}

func (handler *tweets) DeleteTweet(cntx echo.Context) error {
	ctx := cntx.Request().Context()

	tweetID := dto.DeleteTweet{}
	if err := cntx.Bind(&tweetID); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := tweetID.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := handler.app.DeleteTweet(ctx, tweetID.ID); err != nil {
		return err
	}

	return cntx.JSON(http.StatusOK, entity.MessageSuccess{
		Message: "Tweet Deleted Successfully",
	})

}
