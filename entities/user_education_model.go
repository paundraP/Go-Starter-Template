package entities

import (
	"time"

	"github.com/google/uuid"
)

type UserEducation struct {
	ID           uuid.UUID `gorm:"type:uuid;primary_key;not null" json:"id"`
	UserID       uuid.UUID `json:"user_id"`
	SchoolName   string    `json:"school_name"`
	Degree       string    `json:"degree"`
	FieldOfStudy string    `json:"field_of_study"`
	Description  string    `json:"description"`
	StartedAt    time.Time `json:"started_at"`
	EndedAt      time.Time `json:"ended_at"`

	User *User `gorm:"foreignKey:UserID"`
	Timestamp
}
