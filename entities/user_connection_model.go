package entities

import "github.com/google/uuid"

type UserConnection struct {
	UserID          uuid.UUID `gorm:"type:uuid;primary_key" json:"user_id"`
	ConnectedWithID uuid.UUID `gorm:"type:uuid;primary_key" json:"connected_with_id"`

	User       *User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Connection *User `gorm:"foreignKey:ConnectedWithID;constraint:OnDelete:CASCADE"`
	Timestamp
}
