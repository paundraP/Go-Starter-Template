package seeder

import (
	"Go-Starter-Template/internal/utils"
	"Go-Starter-Template/pkg/entities"
	"encoding/json"
	"errors"
	"gorm.io/gorm"
	"log"
	"os"
)

func SeedingUser(db *gorm.DB) error {
	file, err := os.Open("database/seeder/data/user.json")
	if err != nil {
		log.Fatalf("Error opening seed data file: %v", err)
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalf("Error closing seed data file: %v", err)
		}
	}(file)

	var users []entities.User
	if err := json.NewDecoder(file).Decode(&users); err != nil {
		log.Fatalf("Error decoding seed data: %v", err)
		return err
	}

	for i, user := range users {
		var existingUser entities.User
		if err := db.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
			log.Printf("Skipping user %d: email %s already exists", i, user.Email)
			continue
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("Error checking user %d: %v", i, err)
			return err
		}

		hashedPassword, err := utils.HashPassword(user.Password)
		if err != nil {
			log.Fatalf("Error hashing password for user %d: %v", i, err)
			return err
		}
		user.Password = hashedPassword

		if result := db.Create(&user); result.Error != nil {
			log.Printf("Error inserting user %d: %v", i, result.Error)
		} else {
			log.Printf("Inserted user %d: %s", i, user.Email)
		}
	}

	log.Println("seeding user completed successfully!")
	return nil
}
