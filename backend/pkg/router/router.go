package router

import (
	"github.com/ecto0310/online_judge/backend/pkg/users"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRouter() *echo.Echo {
	r := echo.New()

	r.Pre(middleware.RemoveTrailingSlash())

	r.Use(middleware.CORS())
	r.Use(middleware.Logger())
	r.Use(middleware.Recover())

	r.POST("/register", users.Register)
	r.POST("/login", users.Login)
	r.GET("/logout", users.Logout)

	return r
}
