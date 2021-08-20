package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRouter() *echo.Echo {
	r := echo.New()

	r.Pre(middleware.RemoveTrailingSlash())

	r.Use(middleware.CORS())
	r.Use(middleware.Logger())
	r.Use(middleware.Recover())

	r.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	return r
}
