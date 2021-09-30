package users

import (
	"fmt"
	"net/http"

	"github.com/ecto0310/online_judge/backend/pkg/db"
	session "github.com/ipfans/echo-session"
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
	// Get request
	data := LoginData{}
	err := c.Bind(&data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, RegisterResponse{Success: false, Error: "Can't understand your request."})
	}

	// Check user
	hashedPassword := ""
	validUser := true
	user := User{}
	err = db.Db.QueryRow(fmt.Sprintf("SELECT id, name, hashed_password, role FROM users WHERE name='%s'", data.Name)).Scan(&user.Id, &user.Name, &hashedPassword, &user.Role)
	if err != nil {
		validUser = false
		// dummy password
		hashedPassword = "$2a$10$YvuCJ0AY.IHkE1xjoGZ.VOj8Bkol..cWYE7bgW6ycqaXEHQiP7iZO"
	}
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(data.Password))
	if !validUser || err != nil {
		return c.JSON(http.StatusUnauthorized, LoginResponse{Success: false, Error: "Incorrect name or password"})
	}

	// Generate Response
	session := session.Default(c)
	session.Set("id", user.Id)
	session.Set("name", user.Name)
	session.Set("role", user.Role)
	session.Set("login", true)
	session.Save()
	return c.JSON(http.StatusOK, LoginResponse{Success: true, Error: "", User: LoginResponseUser{Name: user.Name, Role: user.Role}})
}

func Logout(c echo.Context) error {
	session := session.Default(c)
	fmt.Println(session.Get("login"))
	session.Set("login", false)
	session.Save()
	return c.JSON(http.StatusOK, LogoutResponse{Success: true, Error: ""})
}
