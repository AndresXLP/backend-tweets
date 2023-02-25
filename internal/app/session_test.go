package app_test

import (
	"context"
	"errors"
	"testing"

	"github.com/andresxlp/backend-twitter/internal/app"
	"github.com/andresxlp/backend-twitter/internal/domain/dto"
	"github.com/andresxlp/backend-twitter/internal/domain/entity"
	mocks "github.com/andresxlp/backend-twitter/mocks/domain/ports/postgres/repo"
	mocks2 "github.com/andresxlp/backend-twitter/mocks/utils"
	"github.com/stretchr/testify/suite"
)

var (
	ctxTrace    = context.Background()
	errExpected = errors.New("error")

	loginData = dto.Login{
		Email:    "test@test.com",
		Password: "123456",
	}
)

type sessionTestSuite struct {
	suite.Suite
	sessionRepo *mocks.Repository
	jwt         *mocks2.JWT
	bcrypt      *mocks2.Bcrypt
	underTest   app.Session
}

func TestSessionSuite(t *testing.T) {
	suite.Run(t, new(sessionTestSuite))
}

func (suite *sessionTestSuite) SetupTest() {
	suite.sessionRepo = &mocks.Repository{}
	suite.jwt = &mocks2.JWT{}
	suite.bcrypt = &mocks2.Bcrypt{}
	suite.underTest = app.NewSessionApp(
		suite.sessionRepo,
		suite.jwt,
		suite.bcrypt,
	)
}

func (suite *sessionTestSuite) TestLogin_WhenGetUserFail() {
	suite.sessionRepo.Mock.On("GetUser", ctxTrace, "test@test.com").
		Return(entity.User{}, errExpected)

	_, err := suite.underTest.Login(ctxTrace, loginData)
	suite.Error(err)
}

func (suite *sessionTestSuite) TestLogin_WhenUserNotExist() {
	suite.sessionRepo.Mock.On("GetUser", ctxTrace, "test@test.com").
		Return(entity.User{}, nil)

	_, err := suite.underTest.Login(ctxTrace, loginData)
	suite.Error(err)
}

func (suite *sessionTestSuite) TestLogin_WhenPasswordNotMatch() {
	suite.sessionRepo.Mock.On("GetUser", ctxTrace, "test@test.com").
		Return(entity.User{ID: 1, Password: []byte("hashedPasswordDB")}, nil)

	suite.bcrypt.Mock.On("ValidatePassword", []byte("hashedPasswordDB"), "123456").
		Return(false)

	_, err := suite.underTest.Login(ctxTrace, loginData)
	suite.Error(err)
}

func (suite *sessionTestSuite) TestLogin_WhenGenerateTokenFail() {
	suite.sessionRepo.Mock.On("GetUser", ctxTrace, "test@test.com").
		Return(entity.User{ID: 1, Email: "test@test.com", Password: []byte("hashedPasswordDB")}, nil)

	suite.bcrypt.Mock.On("ValidatePassword", []byte("hashedPasswordDB"), "123456").
		Return(true)

	suite.jwt.Mock.On("GenerateToken", "test@test.com").
		Return("", errExpected)

	_, err := suite.underTest.Login(ctxTrace, loginData)
	suite.Error(err)
}

func (suite *sessionTestSuite) TestLogin_WhenSuccess() {
	suite.sessionRepo.Mock.On("GetUser", ctxTrace, "test@test.com").
		Return(entity.User{ID: 1, Email: "test@test.com", Password: []byte("hashedPasswordDB")}, nil)

	suite.bcrypt.Mock.On("ValidatePassword", []byte("hashedPasswordDB"), "123456").
		Return(true)

	suite.jwt.Mock.On("GenerateToken", "test@test.com").
		Return("TokenJWTExpected", nil)

	token, err := suite.underTest.Login(ctxTrace, loginData)
	suite.NoError(err)
	suite.Equal(token, "TokenJWTExpected")
}
