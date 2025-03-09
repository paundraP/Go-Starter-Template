package entities

import (
	"github.com/google/uuid"
)

type User struct {
	ID             uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Name           string    `json:"name"`
	Password       string    `json:"password"`
	Email          string    `json:"email"`
	About          string    `json:"about"`
	Address        string    `json:"address"`
	CurrentTitle   string    `json:"current_title"`
	ProfilePicture string    `json:"profile_picture"`
	Headline       string    `json:"headline"`
	IsPremium      bool      `json:"is_premium"`
	Role           string    `json:"role"`
	Timestamp
}
