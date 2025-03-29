package main

import (
	"Go-Starter-Template/cmd/config"
	"Go-Starter-Template/internal/utils"
)

func main() {
	utils.LoadEnv()
	db, err := config.ConnectDB()
	if err != nil {
		panic(err)
	}

	app, err := config.NewApp(db)
	if err != nil {
		panic(err)
	}

	err = app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
