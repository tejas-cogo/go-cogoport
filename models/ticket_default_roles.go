package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TicketDefaultRole struct {
	gorm.Model
	PerformedByID       uuid.UUID `gorm:"type:uuid"`
	TicketDefaultTypeID uint      `gorm:"not null"`
	RoleID              uuid.UUID `gorm:"type:uuid ; not null"`
	UserID              uuid.UUID `gorm:"type:uuid"`
	Level               int       `gorm:"not null;default:3"`
	ClosureAuthorizer   uuid.UUID `gorm:"type:uuid"`
	Status              string    `gorm:"not null;default:'active'"`
}
