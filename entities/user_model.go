package entities

import (
	"github.com/google/uuid"
)

type User struct {
	ID             uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Name           string    `json:"name"`
	Username       string    `json:"username"`
	Email          string    `json:"email"`
	Password       string    `json:"password"`
	Subscribe      bool      `gorm:"default:false" json:"subscribe"`
	Contact        string    `json:"contact"`
	ProfilePicture string    `json:"profile_picture"`
	Role           string    `json:"role"`
	Verified       bool      `gorm:"default:false" json:"verified"`
	ActivePoint    int       `gorm:"default:0" json:"active_point"`
	LevelPoint     int       `gorm:"default:0" json:"level_point"`

	Timestamp
}
