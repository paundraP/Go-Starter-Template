package migration

import (
	"Go-Starter-Template/entities"
	"fmt"
	"log"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
	if err := db.AutoMigrate(&entities.User{}); err != nil {
		log.Fatalf("Error migrating user database: %v", err)
		return err
	}
	if err := db.AutoMigrate(&entities.Transaction{}); err != nil {
		log.Fatalf("Error migrating user database: %v", err)
		return err
	}
	if err := db.AutoMigrate(&entities.Companies{}); err != nil {
		log.Fatalf("Error migrating companies database: %v", err)
		return err
	}
	if err := db.AutoMigrate(&entities.JobApplication{}); err != nil {
		log.Fatalf("Error migrating job application database: %v", err)
		return err
	}
	if err := db.AutoMigrate(entities.Job{}); err != nil {
		log.Fatalf("Error migrating job database: %v", err)
		return err
	}
	if err := db.AutoMigrate(&entities.JobSkill{}); err != nil {
		log.Fatalf("Error migrating job skill database: %v", err)
		return err
	}
	if err := db.AutoMigrate(&entities.Post{}); err != nil {
		log.Fatalf("Error migrating post database: %v", err)
		return err
	}
	if err := db.AutoMigrate(&entities.Skill{}); err != nil {
		log.Fatalf("Error migrating skill database: %v", err)
		return err
	}
	if err := db.AutoMigrate(&entities.UserConnection{}); err != nil {
		log.Fatalf("Error migrating user connection database: %v", err)
		return err
	}
	if err := db.AutoMigrate(&entities.UserEducation{}); err != nil {
		log.Fatalf("Error migrating user education database: %v", err)
		return err
	}
	if err := db.AutoMigrate(&entities.UserExperience{}); err != nil {
		log.Fatalf("Error migrating user experience database: %v", err)
		return err
	}
	if err := db.AutoMigrate(&entities.UserSkill{}); err != nil {
		log.Fatalf("Error migrating user skill database: %v", err)
		return err
	}
	if err := db.AutoMigrate(&entities.Views{}); err != nil {
		log.Fatalf("Error migrating views database: %v", err)
		return err
	}
	fmt.Println("Database migration complete")
	return nil
}
