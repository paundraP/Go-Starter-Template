package entities

import (
	"github.com/google/uuid"
	"time"
)

type UserExperience struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Title       string    `json:"title"`
	CompanyID   uuid.UUID `json:"company_id"`
	UserID      uuid.UUID `json:"user_id"`
	Location    string    `json:"location"`
	Description string    `json:"description"`
	StartedAt   time.Time `json:"started_at"`
	EndedAt     time.Time `json:"ended_at"`

	User    *User      `gorm:"foreignKey:UserID"`
	Company *Companies `gorm:"foreignKey:CompanyID"`

	Timestamp
}
