package users

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type RegisterData struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type RegisterResponseUser struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

type RegisterResponse struct {
	Success bool                 `json:"success"`
	Error   string               `json:"error"`
	User    RegisterResponseUser `json:"user"`
}

func (u *Users) Register(c echo.Context) error {
	data := RegisterData{}
	err := c.Bind(&data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, RegisterResponse{Success: false, Error: "unknown error"})
	}

	err = validateRegisterData(data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, RegisterResponse{Success: false, Error: err.Error()})
	}

	err = registerUser(u.DB, data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, RegisterResponse{Success: false, Error: err.Error()})
	}

	user, err := getUserData(u.DB, data.Name)
	if err != nil {
		return c.JSON(http.StatusBadRequest, RegisterResponse{Success: false, Error: err.Error()})
	}

	responseUser := RegisterResponseUser{Id: user.Id, Name: user.Name, Role: user.Role}
	return c.JSON(http.StatusOK, RegisterResponse{Success: true, Error: "", User: responseUser})
}

func validateRegisterData(data RegisterData) error {
	if len(data.Name) < 4 {
		return errors.New("insufficient name length")
	}
	if len(data.Password) < 8 {
		return errors.New("insufficient password length")
	}

	return nil
}

func registerUser(db *sql.DB, data RegisterData) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), 10)
	if err != nil {
		return errors.New("unknown error")
	}

	_, err = db.Exec("INSERT INTO users (email, name, hashed_password) VALUES (?, ?, ?)", data.Email, data.Name, hashedPassword)
	if err != nil {
		return errors.New("failed to register to DB")
	}

	return nil
}
