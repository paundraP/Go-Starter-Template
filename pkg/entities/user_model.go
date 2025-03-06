package entities

import (
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Subscribe bool      `gorm:"default:false" json:"subscribe"`
	Contact   string    `json:"contact"`
	Role      string    `json:"role"`
	Verified  bool      `gorm:"default:false" json:"verified"`

	Timestamp
}
