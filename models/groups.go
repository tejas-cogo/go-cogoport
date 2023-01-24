package models

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Group struct {
	gorm.Model
	Name          string         `gorm:"not null:unique:json:varchar[]"`
	Tags          pq.StringArray `gorm:"type:text[]"`
	Status        string         `gorm:"not null:default:'active'"`
	PerformedByID uuid.UUID      `gorm:"type:uuid"`
}
