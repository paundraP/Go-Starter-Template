package main

import (
	"Go-Starter-Template/internal/config"
	databaseconf "Go-Starter-Template/internal/config/databaseConf"
)

func main() {
	db, err := databaseconf.ConnectDB()
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
