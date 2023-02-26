package handler_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/andresxlp/backend-twitter/internal/domain/dto"
	"github.com/andresxlp/backend-twitter/internal/infra/api/handler"
	mocks "github.com/andresxlp/backend-twitter/mocks/app"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
)

const (
	pathSession = "/api/session/login"
	methodLogin = "Login"
)

var (
	errExpected       = errors.New("error")
	wrongRequestLogin = dto.Login{
		Email:    "test",
		Password: "123456",
	}

	requestLogin = dto.Login{
		Email:    "test@test.com",
		Password: "123456",
	}
)

type sessionTestSuite struct {
	suite.Suite
	app       *mocks.Session
	underTest handler.Session
}

func TestSessionSuite(t *testing.T) {
	suite.Run(t, new(sessionTestSuite))
}

func (suite *sessionTestSuite) SetupTest() {
	suite.app = &mocks.Session{}
	suite.underTest = handler.NewSessionHandler(suite.app)
}

func (suite *sessionTestSuite) TestLogin_WhenBindFail() {
	body, _ := json.Marshal("")
	controller := SetupControllerCase(http.MethodPost, pathSession, bytes.NewBuffer(body))
	controller.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	suite.Error(suite.underTest.Login(controller.context))
}

func (suite *sessionTestSuite) TestLogin_WhenValidateFail() {
	body, _ := json.Marshal(wrongRequestLogin)
	controller := SetupControllerCase(http.MethodPost, pathSession, bytes.NewBuffer(body))
	controller.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	suite.Error(suite.underTest.Login(controller.context))
}

func (suite *sessionTestSuite) TestLogin_WhenFail() {
	body, _ := json.Marshal(requestLogin)
	controller := SetupControllerCase(http.MethodPost, pathSession, bytes.NewBuffer(body))
	controller.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	suite.app.Mock.On(methodLogin, ctxTest, requestLogin).
		Return("", errExpected)

	suite.Error(suite.underTest.Login(controller.context))
}

func (suite *sessionTestSuite) TestLogin_WhenSuccess() {
	body, _ := json.Marshal(requestLogin)
	controller := SetupControllerCase(http.MethodPost, pathSession, bytes.NewBuffer(body))
	controller.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	suite.app.Mock.On(methodLogin, ctxTest, requestLogin).
		Return("Token", nil)

	suite.NoError(suite.underTest.Login(controller.context))
	suite.Equal(http.StatusOK, controller.Res.Code)
}
