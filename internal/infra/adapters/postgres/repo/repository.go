package repo

import (
	"context"
	"math"

	"github.com/andresxlp/backend-twitter/internal/domain/dto"
	"github.com/andresxlp/backend-twitter/internal/domain/entity"
	"github.com/andresxlp/backend-twitter/internal/domain/ports/postgres/repo"
	"github.com/andresxlp/backend-twitter/internal/infra/adapters/postgres/models"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) repo.Repository {
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

func (repo repository) GetUser(ctx context.Context, email string) (entity.User, error) {
	userDb := models.User{}
	err := repo.db.WithContext(ctx).
		Where("email = ?", email).
		Find(&userDb).Error
	if err != nil {
		return entity.User{}, err
	}

	return userDb.ToDomainEntity(), nil
}

func (repo repository) CreateTweet(ctx context.Context, tweetData models.Tweet) error {
	err := repo.db.WithContext(ctx).
		Create(&tweetData).Error
	if err != nil {
		return err
	}

	return nil
}
func (repo repository) GetTweets(ctx context.Context, request dto.TweetsRequest) (dto.Pagination, entity.TweetsWithOwners, error) {
	var (
		tweets models.TweetsWithOwner
		count  int64
	)
	repo.db.WithContext(ctx).Table("tweets").Count(&count)
	limit := request.Paginate.Limit
	page := request.Paginate.Page
	offset := (page - 1) * limit
	pageCount := int(math.Ceil(float64(count) / float64(limit)))

	err := repo.db.WithContext(ctx).Table("tweets t").
		Select("t.id,t.content, t.visible,u.name as created_by").
		Where("t.visible = true AND t.deleted_at is null").
		Joins("left join users u on t.created_by = u.id").
		Limit(limit).Offset(offset).Scan(&tweets).Error
	if err != nil {
		return dto.Pagination{}, entity.TweetsWithOwners{}, err
	}

	return dto.Pagination{
		TotalRows:  count,
		TotalPages: pageCount,
		Page:       page,
		Limit:      limit,
	}, tweets.ToDomainEntitySlice(), nil
}
