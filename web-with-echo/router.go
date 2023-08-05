package main

import (
	"github.com/labstack/echo/v4"
	"github.com/tksgi/go-practices/web-with-echo/handler"
)

func newRouter() *echo.Echo {
	e := echo.New()

	e.GET("/", handler.Home)

	return e
}
