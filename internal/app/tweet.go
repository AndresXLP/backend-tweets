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
}

type tweets struct {
	tweetRepo repo.Repository
}

func NewTweetsApp(tweetRepo repo.Repository) Tweets {
	return &tweets{tweetRepo}
}

func (app *tweets) CreateTweet(ctx context.Context, tweetData dto.Tweet) error {
	userID := ctx.Value("userID").(int)

	var tweetModel models.Tweets

	tweetModel.BuildModel(tweetData, userID)
	if err := app.tweetRepo.CreateTweet(ctx, tweetModel); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return nil
}
