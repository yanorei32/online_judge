package users

import (
	"net/http"

	"github.com/ecto0310/online_judge/backend/pkg/db"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

// Response to request to post /register
func Register(c echo.Context) error {
	// Get request
	data := RegisterData{}
	err := c.Bind(&data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, RegisterResponse{Success: false, Error: "Can't understand your request."})
	}
	if len(data.Password) < 8 {
		return c.JSON(http.StatusBadRequest, RegisterResponse{Success: false, Error: "Password must be at least 8 characters long."})
	}

	// Generate password hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), 10)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, RegisterResponse{Success: false, Error: "Internal error has occurred."})
	}

	// Register to DB
	res, err := db.Db.Exec("INSERT INTO users (name, hashed_password) VALUES (?, ?)", data.Name, hashedPassword)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, RegisterResponse{Success: false, Error: "Internal error has occurred."})
	}

	// Generate Response
	id, err := res.LastInsertId()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, RegisterResponse{Success: false, Error: "Internal error has occurred."})
	}
	user := RegisterResponseUser{id, data.Name}
	return c.JSON(http.StatusOK, RegisterResponse{Success: true, Error: "", User: user})
}

func Login(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func Logout(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
