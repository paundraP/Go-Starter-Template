package entities

import "github.com/google/uuid"

type Post struct {
	ID      uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	UserID  uuid.UUID `json:"user_id"`
	Asset   string    `json:"asset"`
	Content string    `json:"content"`

	User *User `gorm:"foreignKey:UserID"`
	Timestamp
}
