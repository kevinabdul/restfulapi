package routes

import (
	"github.com/labstack/echo/v4"
)

var e *echo.Echo

func New() *echo.Echo {
	e = echo.New()

	registerRootMiddlewares()

	registerUserRoutes()

	registerBookRoutes()

	return e
}