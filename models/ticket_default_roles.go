package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TicketDefaultRole struct {
	gorm.Model
	PerformedByID       uuid.UUID `gorm:"type:uuid"`
	TicketDefaultTypeID uint      `gorm:"not null"`
	RoleID              uuid.UUID `gorm:"type:uuid"`
	UserID              uuid.UUID `gorm:"type:uuid"`
	Level               int       `gorm:"not null"`
	Status              string    `gorm:"not null:default:'active'"`
}
