package main

import (
	"fmt"
	"os"

	"github.com/ecto0310/online_judge/backend/pkg/server"
)

func main() {
	db, err := server.CreateDbConnection(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE")))
	if err != nil {
		panic(err)
	}

	store, err := server.CreateSessionStore(fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")), os.Getenv("REDIS_PASSWORD"))
	if err != nil {
		panic(err)
	}

	s := server.CreateServer(db, store)
	server.AddRouting(db, s)
	s.Logger.Fatal(s.Start(":1323"))
}
