package entities

import "github.com/google/uuid"

type JobSkill struct {
	JobID   uuid.UUID `gorm:"type:uuid;primary_key" json:"job_id"`
	SkillID uuid.UUID `gorm:"type:uuid;primary_key" json:"skill_id"`

	Job   *Job   `gorm:"foreignKey:JobID;constraint:OnDelete:CASCADE"`
	Skill *Skill `gorm:"foreignKey:SkillID;constraint:OnDelete:CASCADE"`
	Timestamp
}
