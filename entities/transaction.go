package entities

import "github.com/google/uuid"

type Transaction struct {
	ID      uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	UserID  uuid.UUID `json:"user_id"`
	Status  string    `json:"status"`
	Invoice string    `json:"invoice"`
	OrderID string    `json:"order_id"`

	User *User `gorm:"foreignKey:UserID"`
	Timestamp
}
