package entities

import "github.com/google/uuid"

type Companies struct {
	ID    uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Name  string    `json:"name"`
	Slug  string    `json:"slug"`
	About string    `json:"about"`

	Timestamp
}
