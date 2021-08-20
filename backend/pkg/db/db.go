package db

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

// Generate connection to the DB.
// Try until a successful connection is established.
func Init() {
	for {
		db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE")))
		if err == nil {
			err = db.Ping()
			if err == nil {
				Db = db
				break
			}
		}
		fmt.Println("Failed to connect to DB, retry after 5 seconds")
		time.Sleep(time.Second * 5)
	}
}
