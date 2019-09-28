package controllers

import (
	"github.com/labstack/echo"
	"net/http"
)

func Profile(cnt echo.Context) error {
	return cnt.String(http.StatusOK, "profile")
}
