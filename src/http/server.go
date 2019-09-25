package http

import (
	"github.com/labstack/echo"
	"net/http"
)

func New() {
	server := echo.New()
	server.HidePort = true
	server.HideBanner = true

	server.GET("/hui", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	server.Start(":1234")

}

func routes() {

}
