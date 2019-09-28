package controllers

import (
	"github.com/labstack/echo"
	"net/http"
)

func Authorization(cnt echo.Context) error {
	return cnt.String(http.StatusOK, "authorization")
}
