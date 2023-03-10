package dto

import (
	"context"

	"github.com/go-playground/mold/v4/modifiers"
)

var conform = modifiers.New()

type Tweets struct {
	ID        int    `json:"id" param:"id" swaggerignore:"true"`
	Content   string `json:"content" validate:"required"`
	CreatedBy string `json:"created_by" swaggerignore:"true"`
	Visible   *bool  `json:"visible" validate:"boolean"`
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

func (t *TweetsRequest) SetDefault() {
	_ = conform.Struct(context.Background(), t)
}

type DeleteTweet struct {
	ID int `param:"id" validate:"required"`
}

func (d *DeleteTweet) Validate() error {
	return validate.Struct(d)
}
