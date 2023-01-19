package models

import (
	gormjsonb "github.com/dariubs/gorm-jsonb"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TicketDefaultType struct {
	gorm.Model
	PerformedByID uuid.UUID `gorm:"type:uuid"`
	TicketType        string          `gorm:"not null"`
	AdditionalOptions gormjsonb.JSONB `gorm:"type:json"`
	Status            string          `gorm:"not null:default:'active'"`
}
