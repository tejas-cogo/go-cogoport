package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TicketSpectator struct {
	gorm.Model
	PerformedByID uuid.UUID `gorm:"type:uuid"`
	TicketID      uint      `gorm:"not null"`
	UserID        string    `gorm:"not null"`
	Status        string    `gorm:"not null:default:'active'"`
}

type SpectatorActivity struct {
	TicketID        uint
	SpectatorUserID uint
	PerformedByID   uuid.UUID `gorm:"type:uuid"`
	Description     string
}
