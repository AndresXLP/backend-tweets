package router

import (
	"github.com/andresxlp/backend-twitter/internal/infra/api/handler"
	"github.com/andresxlp/backend-twitter/internal/infra/api/router/groups"
	"github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"
)

type Router struct {
	server       *echo.Echo
	userGroup    groups.User
	sessionGroup groups.Session
}

func New(server *echo.Echo, userGroup groups.User, sessionGroup groups.Session) *Router {
	return &Router{
		server,
		userGroup,
		sessionGroup,
	}
}

func (r *Router) Init() {
	r.server.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status} latency=${latency_human}\n",
	}))
	r.server.Use(middleware.Recover())

	basePath := r.server.Group("/api")
	basePath.GET("/health", handler.HealthCheck)

	r.userGroup.Resource(basePath)
	r.sessionGroup.Resource(basePath)

}
