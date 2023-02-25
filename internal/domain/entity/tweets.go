package entity

import "github.com/andresxlp/backend-twitter/internal/domain/dto"

type Tweets struct {
	ID        int
	Content   string
	CreatedBy string
	Visible   *bool
}

type TweetsWithOwners []Tweets

func (t *TweetsWithOwners) Add(tweets Tweets) {
	*t = append(*t, tweets)
}

func (t *Tweets) ToDomainDTOSingle() dto.Tweets {
	return dto.Tweets{
		ID:        t.ID,
		Content:   t.Content,
		CreatedBy: t.CreatedBy,
		Visible:   t.Visible,
	}
}

func (t *TweetsWithOwners) ToDomainDTOSlice() dto.TweetsWithOwner {
	var tweets dto.TweetsWithOwner

	for _, itme := range *t {
		tweets.Add(itme.ToDomainDTOSingle())
	}

	return tweets
}
