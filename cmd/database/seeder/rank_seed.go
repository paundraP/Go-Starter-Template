package seeder

import (
	"Go-Starter-Template/entities"
	"encoding/json"
	"gorm.io/gorm"
	"log"
	"os"
)

func SeedingRank(db *gorm.DB) error {
	file, err := os.Open("cmd/database/seeder/data/rank.json")
	if err != nil {
		log.Fatalf("error opening file: %v", err)
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalf("Error closing seed data file: %v", err)
		}
	}(file)

	var ranks []entities.Rank
	if err := json.NewDecoder(file).Decode(&ranks); err != nil {
		log.Fatalf("error decoding seed data: %v", err)
		return err
	}

	for _, rank := range ranks {
		if err := db.Create(&rank).Error; err != nil {
			log.Fatalf("error seeding data: %v", err)
		} else {
			log.Printf("seeded rank: %v", rank.Name)
		}
	}
	log.Println("seeding rank completed successfully!")
	return nil
}
