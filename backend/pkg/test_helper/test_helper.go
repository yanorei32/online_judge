package test_helper

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/alicebob/miniredis/v2"
	"github.com/ecto0310/online_judge/backend/pkg/server"
	"github.com/labstack/echo/v4"
)

func CreateTestServer() (*echo.Echo, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}
	redis, err := miniredis.Run()
	if err != nil {
		return nil, nil, err
	}
	store, err := server.CreateSessionStore(redis.Addr(), "")
	if err != nil {
		return nil, nil, err
	}
	e := server.CreateServer(db, store)
	server.AddRouting(db, e)

	return e, mock, err
}
