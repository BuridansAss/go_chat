package http

import (
	"buridansAss/chat/src/controllers"
	"github.com/labstack/echo"
)

type Api struct {
	Port string
}

func New() {
	server := echo.New()
	server.HidePort = true
	server.HideBanner = true

	server.POST("/authorization", controllers.Authorization)
	server.POST("/registration", controllers.Registration)
	server.GET("/profile", controllers.Profile)

	server.Start(":1234")
}
