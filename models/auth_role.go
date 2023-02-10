package models

import (
	"github.com/google/uuid"
)

type AuthRole struct {
	ID            uuid.UUID
	Name          string
	StakeholderId uuid.UUID
	Status        string
}

type AuthRoleData struct {
	ID            uuid.UUID
	Name          string    `json:"name"`
	StakeholderId uuid.UUID `gorm:"type:uuid"`
	Status        string    `gorm:"not null:default:'active'"`
}
