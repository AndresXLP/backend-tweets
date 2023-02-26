package main

import (
	"fmt"
	"log"

	"github.com/andresxlp/backend-twitter/cmd/providers"
	"github.com/andresxlp/backend-twitter/config"
	"github.com/andresxlp/backend-twitter/internal/infra/api/router"
	"github.com/labstack/echo/v4"
)

var (
	serverHost = config.Environments().ServerHost
	serverPort = config.Environments().ServerPort
)

// main
//	@title			Tweet Backend
//	@version		1.0.0
//	@description	Backend Basic Tweets with Authentication
//	@license.name	Andres Puello
//	@BasePath		/api
//	@schemes		http
func main() {
	container := providers.BuildContainer()

	err := container.Invoke(func(router *router.Router, server *echo.Echo) {
		router.Init()
		server.Logger.Fatal(server.Start(fmt.Sprintf("%s:%d", serverHost, serverPort)))
	})

	if err != nil {
		log.Panic(err)
	}
}
