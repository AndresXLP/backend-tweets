package app_test

import (
	"testing"

	"github.com/andresxlp/backend-twitter/internal/app"
	"github.com/andresxlp/backend-twitter/internal/domain/dto"
	"github.com/andresxlp/backend-twitter/internal/domain/entity"
	"github.com/andresxlp/backend-twitter/internal/infra/adapters/postgres/models"
	mocks "github.com/andresxlp/backend-twitter/mocks/domain/ports/postgres/repo"
	"github.com/stretchr/testify/suite"
)

type userTestSuite struct {
	suite.Suite
	userRepo  *mocks.Repository
	underTest app.User
}

func TestUserSuite(t *testing.T) {
	suite.Run(t, new(userTestSuite))
}

func (suite *userTestSuite) SetupTest() {
	suite.userRepo = &mocks.Repository{}
	suite.underTest = app.NewUserApp(suite.userRepo)
}

func (suite *userTestSuite) TestCreateUser_WhenGetUserFail() {
	suite.userRepo.Mock.On("GetUser", ctxTrace, "test@test.com").
		Return(entity.User{}, errExpected)

	suite.Error(suite.underTest.CreateUser(ctxTrace, dto.NewUser{
		User: dto.User{
			Email: "test@test.com",
		},
	}))
}

func (suite *userTestSuite) TestCreateUser_WhenUserExist() {
	suite.userRepo.Mock.On("GetUser", ctxTrace, "test@test.com").
		Return(entity.User{ID: 1}, nil)

	suite.Error(suite.underTest.CreateUser(ctxTrace, dto.NewUser{
		User: dto.User{
			Email: "test@test.com",
		},
	}))
}

func (suite *userTestSuite) TestCreateUser_WhenFail() {
	suite.userRepo.Mock.On("GetUser", ctxTrace, "test@test.com").
		Return(entity.User{ID: 0}, nil)

	suite.userRepo.Mock.On("CreateUser", ctxTrace, models.User{
		Email:    "test@test.com",
		Password: []byte("123"),
	}).Return(errExpected)

	suite.Error(suite.underTest.CreateUser(ctxTrace, dto.NewUser{
		User: dto.User{
			Email: "test@test.com",
		},
		Password: "123",
	}))
}

func (suite *userTestSuite) TestCreateUser_WhenSuccess() {
	suite.userRepo.Mock.On("GetUser", ctxTrace, "test@test.com").
		Return(entity.User{ID: 0}, nil)

	suite.userRepo.Mock.On("CreateUser", ctxTrace, models.User{
		Email:    "test@test.com",
		Password: []byte("123"),
	}).Return(nil)

	suite.NoError(suite.underTest.CreateUser(ctxTrace, dto.NewUser{
		User: dto.User{
			Email: "test@test.com",
		},
		Password: "123",
	}))
}
