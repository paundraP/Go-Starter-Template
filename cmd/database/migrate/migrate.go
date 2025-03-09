package migration

import (
	entities2 "Go-Starter-Template/entities"
	"fmt"
	"log"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
	if err := db.AutoMigrate(&entities2.User{}); err != nil {
		log.Fatalf("Error migrating user database: %v", err)
		return err
	}
	if err := db.AutoMigrate(&entities2.Transaction{}); err != nil {
		log.Fatalf("Error migrating user database: %v", err)
		return err
	}
	fmt.Println("Database migration complete")
	return nil
}
