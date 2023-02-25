package app

import (
	"context"
	"net/http"

	"github.com/andresxlp/backend-twitter/internal/domain/dto"
	"github.com/andresxlp/backend-twitter/internal/domain/ports/postgres/repo"
	"github.com/andresxlp/backend-twitter/internal/infra/adapters/postgres/models"
	"github.com/labstack/echo/v4"
)

const (
	TweetNotFound = "tweet not found"
)

type Tweets interface {
	CreateTweet(ctx context.Context, tweetData dto.Tweets) error
	GetTweets(ctx context.Context, request dto.TweetsRequest) (dto.Pagination, error)
	UpdateTweet(ctx context.Context, updateData dto.Tweets) error
	GetTweetByID(ctx context.Context, idTweet int) (dto.Tweets, error)
	GetTweetByIDAndUserID(ctx context.Context, idTweet, userID int) (dto.Tweets, error)
	DeleteTweet(ctx context.Context, tweetID int) error
}

type tweets struct {
	tweetRepo repo.Repository
}

func NewTweetsApp(tweetRepo repo.Repository) Tweets {
	return &tweets{tweetRepo}
}

func (app *tweets) CreateTweet(ctx context.Context, tweetData dto.Tweets) error {
	userID := ctx.Value("userID").(int)

	var tweetModel models.Tweet

	tweetModel.BuildModel(tweetData, userID)
	if err := app.tweetRepo.CreateTweet(ctx, tweetModel); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func (app *tweets) GetTweets(ctx context.Context, request dto.TweetsRequest) (dto.Pagination, error) {
	pagination, entityTweets, err := app.tweetRepo.GetTweets(ctx, request)
	if err != nil {
		return dto.Pagination{}, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	pagination.Rows = entityTweets.ToDomainDTOSlice()

	return pagination, nil
}

func (app *tweets) GetTweetByID(ctx context.Context, idTweet int) (dto.Tweets, error) {
	entityTweet, err := app.tweetRepo.GetTweetByID(ctx, idTweet)
	if err != nil {
		return dto.Tweets{}, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if entityTweet.ID == 0 {
		return dto.Tweets{}, echo.NewHTTPError(http.StatusNotFound, TweetNotFound)
	}

	return entityTweet.ToDomainDTOSingle(), nil
}

func (app *tweets) GetTweetByIDAndUserID(ctx context.Context, idTweet, userID int) (dto.Tweets, error) {
	entityTweet, err := app.tweetRepo.GetTweetByIDAndUserID(ctx, idTweet, userID)
	if err != nil {
		return dto.Tweets{}, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if entityTweet.ID == 0 {
		return dto.Tweets{}, echo.NewHTTPError(http.StatusNotFound, TweetNotFound)
	}

	return entityTweet.ToDomainDTOSingle(), nil
}

func (app *tweets) UpdateTweet(ctx context.Context, updateData dto.Tweets) error {
	userID := ctx.Value("userID").(int)

	originalTweet, err := app.GetTweetByIDAndUserID(ctx, updateData.ID, userID)
	if err != nil {
		return err
	}

	if originalTweet.ID == 0 {
		return echo.NewHTTPError(http.StatusNotFound, TweetNotFound)
	}

	originalTweet.Content = updateData.Content

	originalTweet.Visible = updateData.Visible

	var modelTweet models.Tweet
	modelTweet.BuildModel(originalTweet, userID)
	if err = app.tweetRepo.UpdateTweet(ctx, modelTweet); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func (app *tweets) DeleteTweet(ctx context.Context, tweetID int) error {
	userID := ctx.Value("userID").(int)
	entityTweet, err := app.tweetRepo.GetTweetByIDAndUserID(ctx, tweetID, userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if entityTweet.ID == 0 {
		return echo.NewHTTPError(http.StatusNotFound, TweetNotFound)
	}

	var deleteTweet models.Tweet
	deleteTweet.BuildModel(entityTweet.ToDomainDTOSingle(), userID)
	if err = app.tweetRepo.DeleteTweet(ctx, deleteTweet); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return nil
}
