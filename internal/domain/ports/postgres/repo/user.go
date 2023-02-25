package repo

import (
	"context"

	"github.com/andresxlp/backend-twitter/internal/domain/entity"
	"github.com/andresxlp/backend-twitter/internal/infra/adapters/postgres/models"
)

type User interface {
	CreateUser(ctx context.Context, newUser models.User) error
	GetUser(ctx context.Context, email string) (entity.User, error)
}
