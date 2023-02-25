package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/andresxlp/backend-twitter/internal/domain/dto"
	"github.com/andresxlp/backend-twitter/internal/infra/api/handler"
	mocks "github.com/andresxlp/backend-twitter/mocks/app"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
)

const (
	pathTweets        = "/api/tweets"
	methodCreateTweet = "CreateTweet"
)

var (
	wrongRequestNewTweet = dto.Tweets{
		ID:        0,
		Content:   "",
		CreatedBy: "",
		Visible:   false,
	}

	requestNewTweet = dto.Tweets{
		ID:        0,
		Content:   "Test",
		CreatedBy: "",
		Visible:   true,
	}
)

type tweetsSuiteTest struct {
	suite.Suite
	app       *mocks.Tweets
	underTest handler.Tweets
}

func TestTweetSuite(t *testing.T) {
	suite.Run(t, new(tweetsSuiteTest))
}

func (suite *tweetsSuiteTest) SetupTest() {
	suite.app = &mocks.Tweets{}
	suite.underTest = handler.NewTweetsHandler(suite.app)
}

func (suite *tweetsSuiteTest) TestCreateTweet_WhenBindFail() {
	body, _ := json.Marshal("")
	controller := SetupControllerCase(http.MethodPost, pathTweets, bytes.NewBuffer(body))
	controller.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	suite.Error(suite.underTest.CreateTweet(controller.context))
}

func (suite *tweetsSuiteTest) TestCreateTweet_WhenValidateFail() {
	body, _ := json.Marshal(wrongRequestNewTweet)
	controller := SetupControllerCase(http.MethodPost, pathTweets, bytes.NewBuffer(body))
	controller.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	suite.Error(suite.underTest.CreateTweet(controller.context))
}

func (suite *tweetsSuiteTest) TestCreateTweet_WhenTokenNoProvided() {
	body, _ := json.Marshal(requestNewTweet)
	controller := SetupControllerCase(http.MethodPost, pathTweets, bytes.NewBuffer(body))
	controller.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	suite.app.Mock.On(methodCreateTweet, ctxTest, requestNewTweet).
		Return(errExpected)

	suite.Error(suite.underTest.CreateTweet(controller.context))
}

func (suite *tweetsSuiteTest) TestCreateTweet_WhenSuccess() {
	body, _ := json.Marshal(requestNewTweet)
	controller := SetupControllerCase(http.MethodPost, pathTweets, bytes.NewBuffer(body))
	controller.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	suite.app.Mock.On(methodCreateTweet, ctxTest, requestNewTweet).
		Return(nil)

	suite.NoError(suite.underTest.CreateTweet(controller.context))
}
