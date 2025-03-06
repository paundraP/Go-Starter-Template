package main

import (
	migration "Go-Starter-Template/cmd/database/migrate"
	"Go-Starter-Template/cmd/database/seeder"
	databaseconf "Go-Starter-Template/internal/config/database_config"
	"flag"
	"fmt"
	"log"

	"gorm.io/gorm"
)

func DatabaseSetUp() (*gorm.DB, error) {
	fmt.Println("Hi!")
	// setting up database (migration and data)
	db, err := databaseconf.ConnectDB()
	if db == nil || err != nil {
		return nil, err
	}

	migrateFlag := flag.Bool("migrate", false, "migrating the database")
	seedFlag := flag.Bool("seed", false, "seeding the database")

	flag.Parse()

	if *migrateFlag {
		if err := migration.Migrate(db); err != nil {
			return nil, err
		}
	}
	if *seedFlag {
		if err := seeder.Seed(db); err != nil {
			return nil, err
		}
	}
	return db, nil
}

func main() {
	db, err := DatabaseSetUp()
	if err != nil {
		log.Fatalf("Error setting up database : %v", err)
	}
	fmt.Println("Database setup successful:", db)
}
