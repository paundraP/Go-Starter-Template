package database

import (
	"Go-Starter-Template/pkg/entities"
	"fmt"
	"gorm.io/gorm"
	"log"
)

func Migrate(db *gorm.DB) error {
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
	if err := db.AutoMigrate(&entities.User{}); err != nil {
		log.Fatalf("Error migrating user database: %v", err)
		return err
	}
	if err := db.AutoMigrate(&entities.Rank{}); err != nil {
		log.Fatalf("Error migrating user database: %v", err)
		return err
	}
	if err := db.AutoMigrate(&entities.Transaction{}); err != nil {
		log.Fatalf("Error migrating user database: %v", err)
		return err
	}
	fmt.Println("Database migration complete")
	return nil
}
