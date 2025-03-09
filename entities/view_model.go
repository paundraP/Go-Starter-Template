package entities

import "github.com/google/uuid"

type Views struct {
	UserID   uuid.UUID `gorm:"type:uuid;primary_key" json:"user_id"`
	ViewerID uuid.UUID `gorm:"type:uuid;primary_key" json:"viewer_id"`

	User   *User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Viewer *User `gorm:"foreignKey:ViewerID;constraint:OnDelete:CASCADE"`

	Timestamp
}
