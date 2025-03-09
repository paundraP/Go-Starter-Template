package main

import (
	migration "Go-Starter-Template/cmd/database/migrate"
	databaseconf "Go-Starter-Template/internal/config/database_config"
	"Go-Starter-Template/internal/utils"
	"flag"
	"fmt"
	"log"

	"gorm.io/gorm"
)

func DatabaseSetUp() (*gorm.DB, error) {
	utils.LoadEnv()
	fmt.Println("Hi!")
	// setting up database (migration and data)
	db, err := databaseconf.ConnectDB()
	if db == nil || err != nil {
		return nil, err
	}

	migrateFlag := flag.Bool("migrate", false, "migrating the database")

	flag.Parse()

	if *migrateFlag {
		if err := migration.Migrate(db); err != nil {
			return nil, err
		}
	}
	return db, nil
}

func main() {
	_, err := DatabaseSetUp()
	if err != nil {
		log.Fatalf("Error setting up database : %v", err)
	}
}
