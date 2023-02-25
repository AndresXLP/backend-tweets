package providers

import (
	"github.com/andresxlp/backend-twitter/internal/app"
	"github.com/andresxlp/backend-twitter/internal/infra/adapters/postgres/repo"
	"github.com/andresxlp/backend-twitter/internal/infra/api/handler"
	"github.com/andresxlp/backend-twitter/internal/infra/api/router"
	"github.com/andresxlp/backend-twitter/internal/infra/api/router/groups"
	"github.com/andresxlp/backend-twitter/internal/infra/resources/postgres"
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
)

var Container *dig.Container

func BuildContainer() *dig.Container {
	Container = dig.New()

	_ = Container.Provide(func() *echo.Echo {
		return echo.New()
	})

	_ = Container.Provide(postgres.NewConnection)

	_ = Container.Provide(router.New)

	_ = Container.Provide(groups.NewUserGroup)
	_ = Container.Provide(groups.NewSessionGroup)

	_ = Container.Provide(handler.NewUserHandler)
	_ = Container.Provide(handler.NewSessionHandler)

	_ = Container.Provide(app.NewUserApp)
	_ = Container.Provide(app.NewSessionApp)

	_ = Container.Provide(repo.NewRepository)

	return Container
}
