package entities

import (
	"gorm.io/gorm"
	"time"
)

type Timestamp struct {
	CreatedAt time.Time `gorm:"type:timestamp" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp" json:"updated_at"`
	DeletedAt gorm.DeletedAt
}
