package seeder

import (
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) error {
	if err := SeedingUser(db); err != nil {
		return err
	}
	if err := SeedingRank(db); err != nil {
		return err
	}
	return nil
}
