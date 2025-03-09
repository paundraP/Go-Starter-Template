package entities

import "github.com/google/uuid"

type UserSkill struct {
	ID      uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	UserID  uuid.UUID `json:"user_id"`
	SkillID uuid.UUID `json:"skill_id"`

	User  *User  `gorm:"foreignKey:UserID"`
	Skill *Skill `gorm:"foreignKey:SkillID"`
	Timestamp
}
