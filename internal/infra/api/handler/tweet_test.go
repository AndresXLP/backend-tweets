package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
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
	methodGetTweets   = "GetTweets"
	methodUpdateTweet = "UpdateTweet"
	methodDeleteTweet = "DeleteTweet"
)

var (
	visibility        = true
	wrongRequestTweet = dto.Tweets{
		ID:        0,
		Content:   "",
		CreatedBy: "",
		Visible:   &visibility,
	}

	requestTweet = dto.Tweets{
		ID:        0,
		Content:   "Test",
		CreatedBy: "",
		Visible:   &visibility,
	}
)

type tweetsTestSuite struct {
	suite.Suite
	app       *mocks.Tweets
	underTest handler.Tweets
}

func TestTweetSuite(t *testing.T) {
	suite.Run(t, new(tweetsTestSuite))
}

func (suite *tweetsTestSuite) SetupTest() {
	suite.app = &mocks.Tweets{}
	suite.underTest = handler.NewTweetsHandler(suite.app)
}

func (suite *tweetsTestSuite) TestCreateTweet_WhenBindFail() {
	body, _ := json.Marshal("")
	controller := SetupControllerCase(http.MethodPost, pathTweets, bytes.NewBuffer(body))
	controller.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	suite.Error(suite.underTest.CreateTweet(controller.context))
}

func (suite *tweetsTestSuite) TestCreateTweet_WhenValidateFail() {
	body, _ := json.Marshal(wrongRequestTweet)
	controller := SetupControllerCase(http.MethodPost, pathTweets, bytes.NewBuffer(body))
	controller.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	suite.Error(suite.underTest.CreateTweet(controller.context))
}

func (suite *tweetsTestSuite) TestCreateTweet_WhenTokenNoProvided() {
	body, _ := json.Marshal(requestTweet)
	controller := SetupControllerCase(http.MethodPost, pathTweets, bytes.NewBuffer(body))
	controller.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	suite.app.Mock.On(methodCreateTweet, ctxTest, requestTweet).
		Return(errExpected)

	suite.Error(suite.underTest.CreateTweet(controller.context))
}

func (suite *tweetsTestSuite) TestCreateTweet_WhenSuccess() {
	body, _ := json.Marshal(requestTweet)
	controller := SetupControllerCase(http.MethodPost, pathTweets, bytes.NewBuffer(body))
	controller.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	suite.app.Mock.On(methodCreateTweet, ctxTest, requestTweet).
		Return(nil)

	suite.NoError(suite.underTest.CreateTweet(controller.context))
}

func (suite *tweetsTestSuite) TestGetTweets_WhenBindFail() {
	q := make(url.Values)
	q.Set("page", "1A")
	controller := SetupControllerCase(http.MethodGet, pathTweets+"/?"+q.Encode(), nil)
	controller.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	suite.Error(suite.underTest.GetTweets(controller.context))
}

func (suite *tweetsTestSuite) TestGetTweets_WhenFail() {
	paginate := dto.Paginate{
		Limit: 10,
		Page:  1,
	}

	request := dto.TweetsRequest{Paginate: paginate}

	q := make(url.Values)
	q.Set("page", "1")
	q.Set("limit", "10")

	controller := SetupControllerCase(http.MethodGet, pathTweets+"/?"+q.Encode(), nil)
	controller.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	suite.app.Mock.On(methodGetTweets, ctxTest, request).
		Return(dto.Pagination{}, errExpected)

	suite.Error(suite.underTest.GetTweets(controller.context))
}

func (suite *tweetsTestSuite) TestGetTweets_WhenSuccess() {
	paginate := dto.Paginate{
		Limit: 10,
		Page:  1,
	}

	request := dto.TweetsRequest{Paginate: paginate}

	q := make(url.Values)
	q.Set("page", "1")
	q.Set("limit", "10")

	controller := SetupControllerCase(http.MethodGet, pathTweets+"/?"+q.Encode(), nil)
	controller.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	suite.app.Mock.On(methodGetTweets, ctxTest, request).
		Return(dto.Pagination{}, nil)

	suite.NoError(suite.underTest.GetTweets(controller.context))
}

func (suite *tweetsTestSuite) TestUpdateTweet_WhenBindFail() {
	body, _ := json.Marshal("")

	controller := SetupControllerCase(http.MethodPut, pathTweets, bytes.NewBuffer(body))
	controller.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	suite.Error(suite.underTest.UpdateTweet(controller.context))
}

func (suite *tweetsTestSuite) TestUpdateTweet_WhenValidateFail() {
	body, _ := json.Marshal(wrongRequestTweet)

	controller := SetupControllerCase(http.MethodPut, pathTweets, bytes.NewBuffer(body))
	controller.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	suite.Error(suite.underTest.UpdateTweet(controller.context))
}

func (suite *tweetsTestSuite) TestUpdateTweet_WhenFail() {
	body, _ := json.Marshal(requestTweet)

	controller := SetupControllerCase(http.MethodPut, pathTweets, bytes.NewBuffer(body))
	controller.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	suite.app.Mock.On(methodUpdateTweet, ctxTest, requestTweet).
		Return(errExpected)

	suite.Error(suite.underTest.UpdateTweet(controller.context))
}

func (suite *tweetsTestSuite) TestUpdateTweet_WhenSuccess() {
	body, _ := json.Marshal(requestTweet)

	controller := SetupControllerCase(http.MethodPut, pathTweets, bytes.NewBuffer(body))
	controller.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	suite.app.Mock.On(methodUpdateTweet, ctxTest, requestTweet).
		Return(nil)

	suite.NoError(suite.underTest.UpdateTweet(controller.context))
}

func (suite *tweetsTestSuite) TestDeleteTweet_WhenBindFail() {
	controller := SetupControllerCase(http.MethodDelete, pathTweets, nil)
	controller.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	controller.context.SetParamNames("id")
	controller.context.SetParamValues("A1")

	suite.Error(suite.underTest.DeleteTweet(controller.context))
}

func (suite *tweetsTestSuite) TestDeleteTweet_WhenValidateFail() {
	controller := SetupControllerCase(http.MethodPut, pathTweets, nil)
	controller.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	suite.Error(suite.underTest.DeleteTweet(controller.context))
}

func (suite *tweetsTestSuite) TestDeleteTweet_WhenFail() {
	controller := SetupControllerCase(http.MethodPut, pathTweets, nil)
	controller.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	controller.context.SetParamNames("id")
	controller.context.SetParamValues("1")

	suite.app.Mock.On(methodDeleteTweet, ctxTest, 1).
		Return(errExpected)

	suite.Error(suite.underTest.DeleteTweet(controller.context))
}

func (suite *tweetsTestSuite) TestDeleteTweet_WhenSuccess() {
	controller := SetupControllerCase(http.MethodPut, pathTweets, nil)
	controller.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	controller.context.SetParamNames("id")
	controller.context.SetParamValues("1")

	suite.app.Mock.On(methodDeleteTweet, ctxTest, 1).
		Return(nil)

	suite.NoError(suite.underTest.DeleteTweet(controller.context))
}
