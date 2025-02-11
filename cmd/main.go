package main

import (
	"Go-Starter-Template/internal/config"
)

func main() {
	db, err := DatabaseSetUp()
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
