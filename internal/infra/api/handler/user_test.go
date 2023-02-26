package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/andresxlp/backend-twitter/internal/domain/dto"
	"github.com/andresxlp/backend-twitter/internal/infra/api/handler"
	mocks "github.com/andresxlp/backend-twitter/mocks/app"
	mocks2 "github.com/andresxlp/backend-twitter/mocks/utils"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
)

const (
	pathUser         = "/api/users"
	methodCreateUser = "CreateUser"
)

var (
	requestUser = dto.NewUser{
		User: dto.User{
			Name:     "test",
			LastName: "tester",
			Email:    "test@test.com",
			Address:  "Test 123",
			Gender:   "m",
			Age:      18,
		},
		Password: "123Test456",
	}
)

type userTestSuite struct {
	suite.Suite
	app       *mocks.User
	underTest handler.User
	bcrypt    *mocks2.Bcrypt
}

func TestUserSuite(t *testing.T) {
	suite.Run(t, new(userTestSuite))
}

func (suite *userTestSuite) SetupTest() {
	suite.app = &mocks.User{}
	suite.bcrypt = &mocks2.Bcrypt{}
	suite.underTest = handler.NewUserHandler(
		suite.app,
		suite.bcrypt,
	)
}

func (suite *userTestSuite) TestCreateUser_WhenBindFail() {
	body, _ := json.Marshal("")
	controller := SetupControllerCase(http.MethodPost, pathUser, bytes.NewBuffer(body))
	controller.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	suite.Error(suite.underTest.CreateUser(controller.context))
}

func (suite *userTestSuite) TestCreateUser_WhenValidateFail() {
	wrongRequestUser := dto.NewUser{
		User:     dto.User{},
		Password: "",
	}
	body, _ := json.Marshal(wrongRequestUser)
	controller := SetupControllerCase(http.MethodPost, pathUser, bytes.NewBuffer(body))
	controller.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	suite.Error(suite.underTest.CreateUser(controller.context))
}

func (suite *userTestSuite) TestCreateUser_WhenFail() {
	body, _ := json.Marshal(requestUser)
	controller := SetupControllerCase(http.MethodPost, pathUser, bytes.NewBuffer(body))
	controller.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	suite.app.Mock.On(methodCreateUser, ctxTest, requestUser).
		Return(errExpected)

	suite.bcrypt.Mock.On("HashPassword", &requestUser.Password)

	suite.Error(suite.underTest.CreateUser(controller.context))
}

func (suite *userTestSuite) TestCreateUser_WhenSuccess() {
	body, _ := json.Marshal(requestUser)
	controller := SetupControllerCase(http.MethodPost, pathUser, bytes.NewBuffer(body))
	controller.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	suite.app.Mock.On(methodCreateUser, ctxTest, requestUser).
		Return(nil)

	suite.bcrypt.Mock.On("HashPassword", &requestUser.Password)

	suite.NoError(suite.underTest.CreateUser(controller.context))
}
