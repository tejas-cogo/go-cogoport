package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TicketDefaultRole struct {
	gorm.Model
	PerformedByID       uuid.UUID `gorm:"type:uuid"`
	TicketDefaultTypeID uint      `gorm:"not null"`
	RoleID              uint      `gorm:"not null"`
	UserID              uint      `gorm:"not null"`
	Level               uint      `gorm:"not null"`
	Status              string    `gorm:"not null:default:'active'"`
}
