package groups

import (
	"github.com/andresxlp/backend-twitter/internal/infra/api/handler"
	"github.com/andresxlp/backend-twitter/internal/infra/api/middleware"
	"github.com/labstack/echo/v4"
)

type Tweets interface {
	Resource(c *echo.Group)
}

type tweets struct {
	tweetsHandler  handler.Tweets
	userMiddleware middleware.UserMiddleware
}

func NewTweetsGroup(tweetsHand handler.Tweets, userMiddleware middleware.UserMiddleware) Tweets {
	return &tweets{
		tweetsHand,
		userMiddleware,
	}
}

func (group tweets) Resource(c *echo.Group) {
	groupPath := c.Group("/tweets")
	groupPath.POST("", group.tweetsHandler.CreateTweet, group.userMiddleware.OnlyUsers)
	groupPath.GET("", group.tweetsHandler.GetTweets)
	groupPath.PUT("/:id", group.tweetsHandler.UpdateTweet, group.userMiddleware.OnlyUsers)
	groupPath.DELETE("/:id", group.tweetsHandler.DeleteTweet, group.userMiddleware.OnlyUsers)
}
