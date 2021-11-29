package users

import (
	"net/http"

	echo_session "github.com/ipfans/echo-session"
	"github.com/labstack/echo/v4"
)

type LogoutResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

func (u *Users) Logout(c echo.Context) error {
	session := echo_session.Default(c)
	login := session.Get("login")
	if login == nil || login == false {
		return c.JSON(http.StatusBadRequest, LogoutResponse{Success: false, Error: "Session is not login"})
	}
	session.Set("login", false)
	session.Save()
	return c.JSON(http.StatusOK, LogoutResponse{Success: true, Error: ""})
}
