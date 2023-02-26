package groups

import (
	"github.com/andresxlp/backend-twitter/internal/infra/api/handler"
	"github.com/labstack/echo/v4"
)

type User interface {
	Resource(c *echo.Group)
}

type user struct {
	userHandler handler.User
}

func NewUserGroup(userHand handler.User) User {
	return &user{userHand}
}

func (group user) Resource(c *echo.Group) {
	groupPath := c.Group("/users")
	groupPath.POST("", group.userHandler.CreateUser)
}
