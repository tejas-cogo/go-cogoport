package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	PerformedByID uuid.UUID `gorm:"type:uuid"`
	Name          string `gorm:"not null:unique"`
	Level         uint   `gorm:"not null"`
	Status        string `gorm:"not null:default:'active'"`
}
