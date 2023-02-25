package repo

import (
	"context"

	"github.com/andresxlp/backend-twitter/internal/domain/dto"
	"github.com/andresxlp/backend-twitter/internal/domain/entity"
	"github.com/andresxlp/backend-twitter/internal/infra/adapters/postgres/models"
)

type Repository interface {
	CreateUser(ctx context.Context, newUser models.User) error
	GetUser(ctx context.Context, email string) (entity.User, error)
	CreateTweet(ctx context.Context, tweetData models.Tweet) error
	GetTweets(ctx context.Context, request dto.TweetsRequest) (dto.Pagination, entity.TweetsWithOwners, error)
	GetTweetByID(ctx context.Context, idTweet int) (entity.Tweets, error)
	UpdateTweet(ctx context.Context, tweet models.Tweet) error
	GetTweetByIDAndUserID(ctx context.Context, idTweet, userID int) (entity.Tweets, error)
	DeleteTweet(ctx context.Context, tweet models.Tweet) error
}
