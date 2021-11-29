package server

import (
	"database/sql"

	"github.com/ecto0310/online_judge/backend/pkg/users"
	_ "github.com/go-sql-driver/mysql"
	echo_session "github.com/ipfans/echo-session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func CreateDbConnection(address string) (*sql.DB, error) {
	db, err := sql.Open("mysql", address)
	return db, err
}

func CreateSessionStore(address string, password string) (echo_session.RedisStore, error) {
	store, err := echo_session.NewRedisStore(32, "tcp", address, password, make([]byte, 32))
	if err != nil {
		return nil, err
	}
	store.Options(echo_session.Options{Path: "/", MaxAge: 86400 * 7})
	return store, nil
}

func CreateServer(db *sql.DB, store echo_session.RedisStore) *echo.Echo {
	r := echo.New()

	r.Pre(middleware.RemoveTrailingSlash())

	r.Use(echo_session.Sessions("SESSION", store))
	r.Use(middleware.CORS())
	r.Use(middleware.Logger())
	r.Use(middleware.Recover())

	return r
}

func AddRouting(db *sql.DB, s *echo.Echo) {
	users := &users.Users{DB: db}
	s.POST("/register", users.Register)
	s.POST("/login", users.Login)
	s.GET("/logout", users.Logout)
}
