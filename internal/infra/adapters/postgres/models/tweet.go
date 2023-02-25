package models

import (
	"github.com/andresxlp/backend-twitter/internal/domain/dto"
	"github.com/andresxlp/backend-twitter/internal/domain/entity"
)

type Tweet struct {
	ID        int    `json:"ID"`
	Content   string `json:"content"`
	CreatedBy int    `json:"created_by"`
	Visible   bool   `json:"visible"`
}

type Tweets struct {
	Tweet
	CreatedBy string `json:"createdBy"`
}

type TweetsWithOwner []Tweets

func (t *Tweet) BuildModel(tweet dto.Tweet, userID int) {
	t.Content = tweet.Content
	t.CreatedBy = userID
	t.Visible = true
}

func (t Tweets) ToDomainEntitySingle() entity.Tweets {
	return entity.Tweets{
		ID:        t.ID,
		Content:   t.Content,
		CreatedBy: t.CreatedBy,
		Visible:   t.Visible,
	}
}

func (t TweetsWithOwner) ToDomainEntitySlice() entity.TweetsWithOwners {
	var tweetsWithOwner entity.TweetsWithOwners

	for _, item := range t {
		tweetsWithOwner.Add(item.ToDomainEntitySingle())
	}
	return tweetsWithOwner
}
