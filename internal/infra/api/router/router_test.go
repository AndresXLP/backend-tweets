package router_test

import (
	"net/http"
	"testing"

	"github.com/andresxlp/backend-twitter/internal/infra/api/router"
	"github.com/andresxlp/backend-twitter/internal/infra/api/router/groups"
	mocks "github.com/andresxlp/backend-twitter/mocks/infra/api/handler"
	mocksMiddleware "github.com/andresxlp/backend-twitter/mocks/infra/api/middleware"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
)

var (
	paths = []string{
		"/api/health",
		"/api/users/",
		"/api/session/login",
		"/api/tweets",
		"/api/tweets/:id",
	}
	methods = []string{
		http.MethodGet,
		http.MethodPost,
		http.MethodPut,
		http.MethodDelete,
		http.MethodPatch,
		http.MethodOptions,
		http.MethodConnect,
		http.MethodHead,
	}
)

type RouterTestSuite struct {
	suite.Suite
	server       *echo.Echo
	userGroup    groups.User
	sessionGroup groups.Session
	tweetsGroup  groups.Tweets

	underTest *router.Router
}

func TestRouterSuite(t *testing.T) {
	suite.Run(t, new(RouterTestSuite))
}

func (suite *RouterTestSuite) SetupTest() {
	suite.server = echo.New()
	suite.userGroup = groups.NewUserGroup(&mocks.User{})
	suite.sessionGroup = groups.NewSessionGroup(&mocks.Session{})
	suite.tweetsGroup = groups.NewTweetsGroup(&mocks.Tweets{}, &mocksMiddleware.UserMiddleware{})
	suite.underTest = router.New(
		suite.server,
		suite.userGroup,
		suite.sessionGroup,
		suite.tweetsGroup,
	)
}

func (suite *RouterTestSuite) TestInit() {
	suite.underTest.Init()

	for _, route := range suite.server.Routes() {
		suite.Contains(paths, route.Path)
		suite.Contains(methods, route.Method)
	}
}
