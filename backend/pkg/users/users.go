package users

import (
	"database/sql"
	"errors"
)

type Users struct {
	DB *sql.DB
}

type User struct {
	Id    int64
	Email string
	Name  string
	Role  string
}

func getUserData(db *sql.DB, condition interface{}) (User, error) {
	data := User{}
	switch condition.(type) {
	case int64:
		err := db.QueryRow("SELECT id, name, role FROM users WHERE id=?", condition).Scan(&data.Id, &data.Name, &data.Role)
		if err != nil {
			return User{}, errors.New("specified user does not exist")
		}
	case string:
		err := db.QueryRow("SELECT id, name, role FROM users WHERE name=?", condition).Scan(&data.Id, &data.Name, &data.Role)
		if err != nil {
			return User{}, errors.New("specified user does not exist")
		}
	default:
		return User{}, errors.New("unknown error")
	}
	return data, nil
}
