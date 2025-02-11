package main

import (
	"Go-Starter-Template/database"
	"Go-Starter-Template/database/seeder"
	"flag"

	"gorm.io/gorm"
)

func DatabaseSetUp() (*gorm.DB, error) {
	// setting up database (migration and data)
	db, err := database.NewDB()
	if db == nil || err != nil {
		return nil, err
	}

	migrate := flag.Bool("migrate", false, "migrating the database")
	seed := flag.Bool("seed", false, "seeding the database")

	flag.Parse()

	if *migrate {
		if err := database.Migrate(db); err != nil {
			return nil, err
		}
	}
	if *seed {
		if err := seeder.Seed(db); err != nil {
			return nil, err
		}
	}
	return db, nil
}
