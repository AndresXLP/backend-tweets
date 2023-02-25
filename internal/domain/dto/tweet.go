package dto

import (
	"context"

	"github.com/go-playground/mold/v4/modifiers"
)

var conform = modifiers.New()

type Tweets struct {
	ID        int    `json:"id" param:"id"`
	Content   string `json:"content" validate:"required"`
	CreatedBy string `json:"created_by"`
	Visible   bool   `json:"visible" validate:"required,boolean"`
}

type TweetsWithOwner []Tweets

func (t *TweetsWithOwner) Add(tweets ...Tweets) {
	*t = append(*t, tweets...)
}

type TweetsRequest struct {
	Paginate Paginate
}

func (t *Tweets) Validate() error {
	return validate.Struct(t)
}

func (t *TweetsRequest) Validate() error {
	_ = conform.Struct(context.Background(), t)
	return validate.Struct(t)
}
