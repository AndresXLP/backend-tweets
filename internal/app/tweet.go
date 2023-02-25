package app

import (
	"context"
	"net/http"

	"github.com/andresxlp/backend-twitter/internal/domain/dto"
	"github.com/andresxlp/backend-twitter/internal/domain/ports/postgres/repo"
	"github.com/andresxlp/backend-twitter/internal/infra/adapters/postgres/models"
	"github.com/labstack/echo/v4"
)

type Tweets interface {
	CreateTweet(ctx context.Context, tweetData dto.Tweet) error
	GetTweets(ctx context.Context, request dto.TweetsRequest) (dto.Pagination, error)
}

type tweets struct {
	tweetRepo repo.Repository
}

func NewTweetsApp(tweetRepo repo.Repository) Tweets {
	return &tweets{tweetRepo}
}

func (app *tweets) CreateTweet(ctx context.Context, tweetData dto.Tweet) error {
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
