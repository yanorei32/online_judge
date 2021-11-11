package users

import (
	"net/http"

	session "github.com/ipfans/echo-session"
	"github.com/labstack/echo/v4"
)

type LogoutResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

func (u *Users) Logout(c echo.Context) error {
	session := session.Default(c)
	session.Set("login", false)
	session.Save()
	return c.JSON(http.StatusOK, LogoutResponse{Success: true, Error: ""})
}
