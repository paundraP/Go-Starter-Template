package entities

import "github.com/google/uuid"

type JobApplication struct {
	ID     uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	UserID uuid.UUID `json:"user_id"`
	JobID  uuid.UUID `json:"job_id"`
	CV     string    `json:"cv"`
	Status string    `json:"status"`

	User *User `gorm:"foreignKey:UserID"`
	Job  *Job  `gorm:"foreignKey:JobID"`
	Timestamp
}
