package app_test

import (
	"context"
	"testing"
	"time"

	"github.com/andresxlp/backend-twitter/internal/app"
	"github.com/andresxlp/backend-twitter/internal/domain/dto"
	"github.com/andresxlp/backend-twitter/internal/domain/entity"
	"github.com/andresxlp/backend-twitter/internal/infra/adapters/postgres/models"
	mocks "github.com/andresxlp/backend-twitter/mocks/domain/ports/postgres/repo"
	"github.com/stretchr/testify/suite"
)

type tweetTestSuite struct {
	suite.Suite
	tweetRepo *mocks.Repository
	underTest app.Tweets
}

func TestTweetSuite(t *testing.T) {
	suite.Run(t, new(tweetTestSuite))
}

func (suite *tweetTestSuite) SetupTest() {
	suite.tweetRepo = &mocks.Repository{}
	suite.underTest = app.NewTweetsApp(suite.tweetRepo)
}

func (suite *tweetTestSuite) TestCreateTweet_WhenFail() {
	ctxTrace = context.WithValue(ctxTrace, "userID", 1)
	suite.tweetRepo.Mock.On("CreateTweet", ctxTrace, models.Tweet{CreatedBy: 1}).
		Return(errExpected)

	suite.Error(suite.underTest.CreateTweet(ctxTrace, dto.Tweets{}))
}

func (suite *tweetTestSuite) TestCreateTweet_WhenSuccess() {
	ctxTrace = context.WithValue(ctxTrace, "userID", 1)
	suite.tweetRepo.Mock.On("CreateTweet", ctxTrace, models.Tweet{CreatedBy: 1}).
		Return(nil)

	suite.NoError(suite.underTest.CreateTweet(ctxTrace, dto.Tweets{}))
}

func (suite *tweetTestSuite) TestGetTweets_WhenFail() {
	suite.tweetRepo.Mock.On("GetTweets", ctxTrace, dto.TweetsRequest{}).
		Return(dto.Pagination{}, entity.TweetsWithOwners{}, errExpected)

	_, err := suite.underTest.GetTweets(ctxTrace, dto.TweetsRequest{})
	suite.Error(err)
}

func (suite *tweetTestSuite) TestGetTweets_WhenGetTweetsSuccess() {
	suite.tweetRepo.Mock.On("GetTweets", ctxTrace, dto.TweetsRequest{}).
		Return(dto.Pagination{}, entity.TweetsWithOwners{}, nil)

	_, err := suite.underTest.GetTweets(ctxTrace, dto.TweetsRequest{})
	suite.NoError(err)
}

func (suite *tweetTestSuite) TestGetTweetByID_WhenFail() {
	suite.tweetRepo.Mock.On("GetTweetByID", ctxTrace, 1).
		Return(entity.Tweets{}, errExpected)

	_, err := suite.underTest.GetTweetByID(ctxTrace, 1)
	suite.Error(err)
}

func (suite *tweetTestSuite) TestGetTweetByID_WhenTweetNotFound() {
	suite.tweetRepo.Mock.On("GetTweetByID", ctxTrace, 1).
		Return(entity.Tweets{}, nil)

	_, err := suite.underTest.GetTweetByID(ctxTrace, 1)
	suite.Error(err)
}

func (suite *tweetTestSuite) TestGetTweetByID_WhenSuccess() {
	suite.tweetRepo.Mock.On("GetTweetByID", ctxTrace, 1).
		Return(entity.Tweets{ID: 1}, nil)

	_, err := suite.underTest.GetTweetByID(ctxTrace, 1)
	suite.NoError(err)
}

func (suite *tweetTestSuite) TestGetTweetByIDAndUserID_WhenFail() {
	suite.tweetRepo.Mock.On("GetTweetByIDAndUserID", ctxTrace, 1, 1).
		Return(entity.Tweets{}, errExpected)

	_, err := suite.underTest.GetTweetByIDAndUserID(ctxTrace, 1, 1)
	suite.Error(err)
}

func (suite *tweetTestSuite) TestGetTweetByIDAndUserID_WhenTweetNotFound() {
	suite.tweetRepo.Mock.On("GetTweetByIDAndUserID", ctxTrace, 1, 1).
		Return(entity.Tweets{}, nil)

	_, err := suite.underTest.GetTweetByIDAndUserID(ctxTrace, 1, 1)
	suite.Error(err)
}

func (suite *tweetTestSuite) TestGetTweetByIDAndUserID_WhenSuccess() {
	suite.tweetRepo.Mock.On("GetTweetByIDAndUserID", ctxTrace, 1, 1).
		Return(entity.Tweets{ID: 1}, nil)

	_, err := suite.underTest.GetTweetByIDAndUserID(ctxTrace, 1, 1)
	suite.NoError(err)
}

func (suite *tweetTestSuite) TestUpdateTweet_WhenGetTweetByIDAndUserID() {
	ctxTrace = context.WithValue(ctxTrace, "userID", 1)

	suite.tweetRepo.Mock.On("GetTweetByIDAndUserID", ctxTrace, 1, 1).
		Return(entity.Tweets{ID: 1}, errExpected)

	suite.Error(suite.underTest.UpdateTweet(ctxTrace, dto.Tweets{ID: 1}))
}

func (suite *tweetTestSuite) TestUpdateTweet_WhenTweetNotFound() {
	ctxTrace = context.WithValue(ctxTrace, "userID", 1)

	suite.tweetRepo.On("GetTweetByIDAndUserID", ctxTrace, 1, 1).
		Return(entity.Tweets{ID: 0}, nil)

	suite.Error(suite.underTest.UpdateTweet(ctxTrace, dto.Tweets{ID: 1}))
}

func (suite *tweetTestSuite) TestUpdateTweet_WhenFail() {
	ctxTrace = context.WithValue(ctxTrace, "userID", 1)

	suite.tweetRepo.On("GetTweetByIDAndUserID", ctxTrace, 1, 1).
		Return(entity.Tweets{ID: 1}, nil)

	suite.tweetRepo.Mock.On("UpdateTweet", ctxTrace, models.Tweet{
		ID:        1,
		Content:   "",
		CreatedBy: 1,
		Visible:   nil,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	}).
		Return(errExpected)

	suite.Error(suite.underTest.UpdateTweet(ctxTrace, dto.Tweets{ID: 1}))
}

func (suite *tweetTestSuite) TestUpdateTweet_WhenSuccess() {
	ctxTrace = context.WithValue(ctxTrace, "userID", 1)

	suite.tweetRepo.On("GetTweetByIDAndUserID", ctxTrace, 1, 1).
		Return(entity.Tweets{ID: 1}, nil)

	suite.tweetRepo.Mock.On("UpdateTweet", ctxTrace, models.Tweet{
		ID:        1,
		CreatedBy: 1,
	}).
		Return(nil)

	suite.NoError(suite.underTest.UpdateTweet(ctxTrace, dto.Tweets{ID: 1}))
}

func (suite *tweetTestSuite) TestDeleteTweet_WhenGetTweetByIDAndUserID() {
	ctxTrace = context.WithValue(ctxTrace, "userID", 1)

	suite.tweetRepo.On("GetTweetByIDAndUserID", ctxTrace, 1, 1).
		Return(entity.Tweets{ID: 0}, errExpected)

	suite.Error(suite.underTest.DeleteTweet(ctxTrace, 1))
}

func (suite *tweetTestSuite) TestDeleteTweet_WhenTweetNotFound() {
	ctxTrace = context.WithValue(ctxTrace, "userID", 1)

	suite.tweetRepo.On("GetTweetByIDAndUserID", ctxTrace, 1, 1).
		Return(entity.Tweets{ID: 0}, nil)

	suite.Error(suite.underTest.DeleteTweet(ctxTrace, 1))
}

func (suite *tweetTestSuite) TestDeleteTweet_WhenFail() {
	ctxTrace = context.WithValue(ctxTrace, "userID", 1)

	suite.tweetRepo.On("GetTweetByIDAndUserID", ctxTrace, 1, 1).
		Return(entity.Tweets{ID: 1}, nil)

	suite.tweetRepo.Mock.On("DeleteTweet", ctxTrace, models.Tweet{
		ID:        1,
		CreatedBy: 1,
	}).
		Return(errExpected)

	suite.Error(suite.underTest.DeleteTweet(ctxTrace, 1))
}

func (suite *tweetTestSuite) TestDeleteTweet_WhenSuccess() {
	ctxTrace = context.WithValue(ctxTrace, "userID", 1)

	suite.tweetRepo.On("GetTweetByIDAndUserID", ctxTrace, 1, 1).
		Return(entity.Tweets{ID: 1}, nil)

	suite.tweetRepo.Mock.On("DeleteTweet", ctxTrace, models.Tweet{
		ID:        1,
		CreatedBy: 1,
	}).
		Return(nil)

	suite.NoError(suite.underTest.DeleteTweet(ctxTrace, 1))
}
