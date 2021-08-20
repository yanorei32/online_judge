package main

import (
	"github.com/ecto0310/online_judge/backend/pkg/db"
	"github.com/ecto0310/online_judge/backend/pkg/router"
)

func main() {
	db.Init()

	r := router.InitRouter()
	r.Logger.Fatal(r.Start(":1323"))
}
