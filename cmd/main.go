package main

import (
	"Go-Starter-Template/cmd/config"
	"Go-Starter-Template/cmd/config/database_config"
	"Go-Starter-Template/internal/utils"
	"log"

	"os"
)

func main() {
	utils.LoadEnv()
	addr := os.Getenv("APP_URL")
	db, err := databaseconf.ConnectDB()
	if err != nil {
		log.Fatalf("error connection to database: %v", err)
	}

	app, err := config.NewApp(db)
	if err != nil {
		log.Fatalf("error config app: %v", err)
	}

	if addr == "" {
		addr = "0.0.0.0:8080"
	}
	err = app.Listen(addr)
	if err != nil {
		log.Fatalf("error starting app: %v", err)
	}
}
