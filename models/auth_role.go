package models

import (
	"github.com/google/uuid"
)

type AuthRole struct {
	ID                  uuid.UUID `gorm:"type:uuid"`
	Name                string
	StakeholderId       uuid.UUID `gorm:"type:uuid"`
	Status              string    `gorm:"not null:default:'active'"`
}
