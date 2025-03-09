package entities

import "github.com/google/uuid"

type Skill struct {
	ID   uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Name string    `json:"name"`

	Jobs []*Job `gorm:"many2many:job_skills" json:"jobs"`
	Timestamp
}
