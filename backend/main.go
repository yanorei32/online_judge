package main

import "github.com/ecto0310/online_judge/backend/pkg/server"

func main() {
	db, err := server.CreateDbConnection()
	if err != nil {
		panic(err)
	}

	s := server.CreateServer(db)
	s.Logger.Fatal(s.Start(":1323"))
}
