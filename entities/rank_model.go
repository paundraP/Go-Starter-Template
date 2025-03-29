package entities

import "github.com/google/uuid"

type Rank struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Name       string    `json:"name"`
	LowerPoint int       `json:"lower_point"`
	UpperPoint int       `json:"upper_point"`

	Timestamp
}
