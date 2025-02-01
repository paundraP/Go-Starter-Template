package main

import "Go-Starter-Template/internal/config"

func main() {
	app, err := config.NewApp()
	if err != nil {
		panic(err)
	}
	err = app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
