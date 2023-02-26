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

//	@Tags			Tweets
//	@Summary		Create Tweet
//	@Description	Create a new Tweet
//	@Produce		json
//	@Param			Authorization	header		string		true	"Token JWT"
//	@Param			request			body		dto.Tweets	true	"Request Body"
//	@Success		201				{object}	entity.Message
//	@Failure		400
//	@Failure		500
//	@Router			/tweets [post]
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

//	@Tags			Tweets
//	@Summary		Get All Tweets
//	@Description	Get all tweets
//	@Produce		json
//	@Param			page	query		string	false	"page to find"
//	@Param			limit	query		string	false	"limit of page"
//	@Success		200		{object}	dto.Pagination
//	@Failure		400
//	@Failure		500
//	@Router			/tweets [get]
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

//	@Tags			Tweets
//	@Summary		Update Tweet
//	@Description	Update a  Tweet
//	@Produce		json
//	@Param			JWT			Authorization	header		string	true	"Bearer Token"
//	@Param			tweet_id	path			int			true	"tweet_id"
//	@Param			request		body			dto.Tweets	true	"request body"
//	@Success		200			{object}		entity.MessageSuccess
//	@Failure		400
//	@Failure		500
//	@Router			/tweets/{tweet_id} [put]
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

//	@Tags			Tweets
//	@Summary		Delete Tweet
//	@Description	Delete a  Tweet
//	@Produce		json
//	@Param			JWT			Authorization	header	string	true	"Bearer Token"
//	@Param			tweet_id	path			string	true	"tweet_id"
//	@Success		200			{object}		entity.MessageSuccess
//	@Failure		400
//	@Failure		500
//	@Router			/tweets/{tweet_id}  [delete]
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
