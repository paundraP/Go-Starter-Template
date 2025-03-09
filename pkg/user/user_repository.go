package user

import (
	"gorm.io/gorm"
)

type (
	UserRepository interface {
	}
	userRepository struct {
		db *gorm.DB
	}
)

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}
