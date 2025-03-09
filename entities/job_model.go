package entities

import "github.com/google/uuid"

type Job struct {
	ID              uuid.UUID `gorm:"type:uuid;primary_key;not null" json:"id"`
	CompanyID       uuid.UUID `json:"company_id"`
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	Location        string    `json:"location"`
	LocationType    string    `json:"location_type"`
	JobType         string    `json:"job_type"`
	ExperienceLevel string    `json:"experience_level"`
	SalaryMin       int       `json:"salary_min"`
	SalaryMax       int       `json:"salary_max"`
	Status          string    `json:"status"`

	Company *Companies `gorm:"foreignKey:CompanyID"`
	Skills  []*Skill   `gorm:"many2many:job_skills" json:"skills"`
	Timestamp
}
