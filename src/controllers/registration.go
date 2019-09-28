package controllers

import (
	"github.com/labstack/echo"
	"net/http"
)

func Registration(cnt echo.Context) error {
	return cnt.String(http.StatusOK, "registration")
}
