package repo

import (
	"context"

	"github.com/andresxlp/backend-twitter/internal/domain/ports/postgres/repo"
	"github.com/andresxlp/backend-twitter/internal/infra/adapters/postgres/models"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) repo.User {
	return repository{db}
}

func (repo repository) CreateUser(ctx context.Context, newUser models.User) error {
	err := repo.db.WithContext(ctx).
		Create(&newUser).Error
	if err != nil {
		return err
	}

	return nil
}
