package dto

import (
	"context"

	"github.com/go-playground/mold/v4/modifiers"
)

var conform = modifiers.New()

type Tweet struct {
	Content string `json:"content" validate:"required,max=255"`
}

type Tweets struct {
	ID        int    `json:"id"`
	Content   string `json:"content"`
	CreatedBy string `json:"created_by"`
	Visible   bool   `json:"visible"`
}

type TweetsWithOwner []Tweets

func (t *TweetsWithOwner) Add(tweets ...Tweets) {
	*t = append(*t, tweets...)
}

type TweetsRequest struct {
	Paginate Paginate
}

func (t *Tweet) Validate() error {
	return validate.Struct(t)
}

func (t *TweetsRequest) Validate() error {
	_ = conform.Struct(context.Background(), t)
	return validate.Struct(t)
}
