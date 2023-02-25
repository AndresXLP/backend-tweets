package models

import "github.com/andresxlp/backend-twitter/internal/domain/dto"

type Tweets struct {
	ID        int
	Content   string
	CreatedBy int
}

func (t *Tweets) BuildModel(tweet dto.Tweet, userID int) {
	t.Content = tweet.Content
	t.CreatedBy = userID
}
