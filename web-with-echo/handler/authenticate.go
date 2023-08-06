package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/tksgi/go-practices/web-with-echo/model"
)

func SignUp(c echo.Context) error {
  user := model.User{
  	Name:      c.FormValue("name"),
  	Password:  c.FormValue("password"),
  }

  err := model.CreateUser(&user)
	if err != nil {
		log.Error(err)
    return c.String(http.StatusBadRequest, fmt.Sprintf("cannot create User: %v", err))
	}
	return c.String(http.StatusOK, fmt.Sprintf("Successfully created User: %+v", &user))
}
