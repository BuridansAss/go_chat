package controllers

import (
	"buridansAss/chat/src/requests"
	"github.com/labstack/echo"
	"net/http"
)

func Registration(cnt echo.Context) error {
	var req requests.UserRequest
	if err := cnt.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user, err := a.accountService.Register(req.Email, req.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return cnt.JSON(http.StatusOK, user)
}
