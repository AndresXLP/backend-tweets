package middleware_test

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/andresxlp/backend-twitter/internal/domain/entity"
	"github.com/andresxlp/backend-twitter/internal/infra/api/middleware"
	mocks "github.com/andresxlp/backend-twitter/mocks/domain/ports/postgres/repo"
	mocks2 "github.com/andresxlp/backend-twitter/mocks/utils"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
)

var (
	errExpected = errors.New("error")
	ctxTest     = context.Background()
)

type MiddlewareCase struct {
	Req     *http.Request
	Res     *httptest.ResponseRecorder
	context echo.Context
}

func SetupMiddlewareCase(method string, url string, body io.Reader) MiddlewareCase {
	e := echo.New()
	req := httptest.NewRequest(method, url, body)
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)

	return MiddlewareCase{req, res, c}
}

func (c MiddlewareCase) Run(cb echo.HandlerFunc, m echo.MiddlewareFunc) error {
	h := m(cb)

	return h(c.context)
}

type userMiddlewareTestSuite struct {
	suite.Suite
	repo      *mocks.Repository
	jwt       *mocks2.JWT
	underTest middleware.UserMiddleware
}

func TestUserMiddlewareSuite(t *testing.T) {
	suite.Run(t, new(userMiddlewareTestSuite))
}

func (suite *userMiddlewareTestSuite) SetupTest() {
	suite.repo = &mocks.Repository{}
	suite.jwt = &mocks2.JWT{}
	suite.underTest = middleware.NewUserMiddleware(suite.repo, suite.jwt)
}

func (suite *userMiddlewareTestSuite) TestOnlyUsers_WhenTokenNoProvide() {
	cb := func(c echo.Context) error {
		return nil
	}

	c := SetupMiddlewareCase(http.MethodPost, "/", nil)
	suite.Error(c.Run(cb, suite.underTest.OnlyUsers))
}

func (suite *userMiddlewareTestSuite) TestOnlyUsers_WhenValidateTokenFail() {
	cb := func(c echo.Context) error {
		return nil
	}
	suite.jwt.Mock.On("ValidateToken", "TokenTester").
		Return("", errExpected)

	c := SetupMiddlewareCase(http.MethodPost, "/", nil)
	c.Req.Header.Set("Authorization", "Bearer TokenTester")

	suite.Error(c.Run(cb, suite.underTest.OnlyUsers))
}

func (suite *userMiddlewareTestSuite) TestOnlyUsers_WhenGetUserFail() {
	cb := func(c echo.Context) error {
		return nil
	}

	suite.jwt.Mock.On("ValidateToken", "TokenTester").
		Return("test@test.com", nil)

	suite.repo.Mock.On("GetUser", ctxTest, "test@test.com").
		Return(entity.User{}, errExpected)

	c := SetupMiddlewareCase(http.MethodPost, "/", nil)
	c.Req.Header.Set("Authorization", "Bearer TokenTester")

	suite.Error(c.Run(cb, suite.underTest.OnlyUsers))
}

func (suite *userMiddlewareTestSuite) TestOnlyUsers_WhenUserNotExist() {
	cb := func(c echo.Context) error {
		return nil
	}

	suite.jwt.Mock.On("ValidateToken", "TokenTester").
		Return("test@test.com", nil)

	suite.repo.Mock.On("GetUser", ctxTest, "test@test.com").
		Return(entity.User{ID: 0}, nil)

	c := SetupMiddlewareCase(http.MethodPost, "/", nil)
	c.Req.Header.Set("Authorization", "Bearer TokenTester")

	suite.Error(c.Run(cb, suite.underTest.OnlyUsers))
}

func (suite *userMiddlewareTestSuite) TestOnlyUsers_WhenSuccess() {
	cb := func(c echo.Context) error {
		return nil
	}

	suite.jwt.Mock.On("ValidateToken", "TokenTester").
		Return("test@test.com", nil)

	suite.repo.Mock.On("GetUser", ctxTest, "test@test.com").
		Return(entity.User{ID: 1}, nil)

	c := SetupMiddlewareCase(http.MethodPost, "/", nil)
	c.Req.Header.Set("Authorization", "Bearer TokenTester")

	suite.NoError(c.Run(cb, suite.underTest.OnlyUsers))
	suite.Equal(c.context.Request().Context().Value("userID").(int), 1)
}
