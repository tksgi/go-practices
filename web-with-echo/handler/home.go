package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/tksgi/go-practices/web-with-echo/model"
)

func Home(c echo.Context) error {
	users, err := model.ListAllUsers()
	if err != nil {
		log.Error(err)
		return c.String(http.StatusBadRequest, "cannot get Users")
	}
	return c.String(http.StatusOK, fmt.Sprintf("%+v", users))
}
