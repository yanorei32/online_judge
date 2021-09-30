package router

import (
	"fmt"
	"os"

	"github.com/ecto0310/online_judge/backend/pkg/users"
	session "github.com/ipfans/echo-session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRouter() *echo.Echo {
	r := echo.New()

	r.Pre(middleware.RemoveTrailingSlash())

	store, err := session.NewRedisStore(32, "tcp", fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")), os.Getenv("REDIS_PASSWORD"), make([]byte, 32))
	if err != nil {
		panic(err)
	}
	store.Options(session.Options{Path: "/", MaxAge: 86400 * 7})
	r.Use(session.Sessions("SESSION", store))

	r.Use(middleware.CORS())
	r.Use(middleware.Logger())
	r.Use(middleware.Recover())

	r.POST("/register", users.Register)
	r.POST("/login", users.Login)
	r.GET("/logout", users.Logout)

	return r
}
