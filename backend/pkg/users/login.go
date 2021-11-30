package users

import (
	"database/sql"
	"errors"
	"net/http"

	echo_session "github.com/ipfans/echo-session"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type LoginData struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type LoginResponseUser struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

type LoginResponse struct {
	Success bool              `json:"success"`
	Error   string            `json:"error"`
	User    LoginResponseUser `json:"user"`
}

func (u *Users) Login(c echo.Context) error {
	data := LoginData{}
	err := c.Bind(&data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, LoginResponse{Success: false, Error: "unknown error"})
	}

	err = checkUser(u.DB, data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, LoginResponse{Success: false, Error: err.Error()})
	}

	user, err := getUserData(u.DB, data.Name)
	if err != nil {
		return c.JSON(http.StatusBadRequest, LoginResponse{Success: false, Error: err.Error()})
	}

	loginUser := LoginResponseUser{Id: user.Id, Name: user.Name, Role: user.Role}
	session := echo_session.Default(c)
	session.Set("id", loginUser.Id)
	session.Set("name", loginUser.Name)
	session.Set("role", loginUser.Role)
	session.Set("login", true)
	session.Save()
	return c.JSON(http.StatusOK, LoginResponse{Success: true, Error: "", User: loginUser})
}

func checkUser(db *sql.DB, data LoginData) error {
	hashedPassword := ""
	err := db.QueryRow("SELECT hashed_password FROM users WHERE name=?", data.Name).Scan(&hashedPassword)
	if err == sql.ErrNoRows {
		return errors.New("specified user does not exist")
	}
	if err != nil {
		return errors.New("failed to get to DB")
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(data.Password))
	if err != nil {
		return errors.New("password is incorrect")
	}

	return nil
}
