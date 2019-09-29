package http

import (
	"buridansAss/chat/src/controllers"
	"buridansAss/chat/src/providers"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

type (
	Api struct {
		logger *logrus.Logger
		config *providers.Config
	}

	/*
	 *любая структура встраивающая fx.In обрабатывается как структура параметров
	 *конструктор ниже будет более читаемым
	 */
	DIApiOptions struct {
		fx.In //сверху описал fx.In

		Config *providers.Config
		Logger *logrus.Logger
	}
)

func New(options DIApiOptions) {

	api := &Api{
		config: options.Config,
		logger: options.Logger,
	}

	server := echo.New()
	server.HidePort = true
	server.HideBanner = true

	server.POST("/authorization", controllers.Authorization)
	server.POST("/registration", controllers.Registration)
	server.GET("/profile", controllers.Profile)

	server.Start(api.config.ListenPort)
}
