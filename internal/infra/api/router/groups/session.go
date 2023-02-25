package groups

import (
	"github.com/andresxlp/backend-twitter/internal/infra/api/handler"
	"github.com/labstack/echo/v4"
)

type Session interface {
	Resource(c *echo.Group)
}

type session struct {
	sessionHandler handler.Session
}

func NewSessionGroup(sessionHand handler.Session) Session {
	return &session{sessionHand}
}

func (group session) Resource(c *echo.Group) {
	groupPath := c.Group("/session")
	groupPath.POST("/login", group.sessionHandler.Login)
}
