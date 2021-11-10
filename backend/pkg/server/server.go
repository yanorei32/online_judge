package server

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	session "github.com/ipfans/echo-session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func CreateDbConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE")))
	return db, err
}

func CreateServer(db *sql.DB) *echo.Echo {
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

	return r
}
