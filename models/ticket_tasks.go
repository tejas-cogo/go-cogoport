package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TicketTask struct {
	gorm.Model
	PerformedByID uuid.UUID `gorm:"type:uuid"`
	TicketID        uint      `gorm:"not null"`
	Title           string    `gorm:"not null"`
	CreatedByUserId uuid.UUID `gorm:"not null"`
	Status          string    `gorm:"not null:default:'active'"`
}
