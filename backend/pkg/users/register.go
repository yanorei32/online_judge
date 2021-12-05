package users

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"github.com/go-playground/validator/v10"

)

var validate *validator.Validate

type RegisterData struct {
	Email    string `json:"email" validate:"required,min=3,max=191"`
	Name     string `json:"name" validate:"required,min=4,max=255,regexp=^[a-zA-Z_]*$"`
	Password string `json:"password" validate:"required,min=8,max=191"`
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

	if errs := validate.Struct(data); errs != nil {
		return c.JSON(
			http.StatusBadRequest,
			RegisterResponse{Success: false, Error: errs.(validator.ValidationErrors).Error()},
		)
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
