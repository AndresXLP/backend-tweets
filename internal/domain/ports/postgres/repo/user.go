package repo

import (
	"context"

	"github.com/andresxlp/backend-twitter/internal/infra/adapters/postgres/models"
)

type User interface {
	CreateUser(ctx context.Context, newUser models.User) error
}
