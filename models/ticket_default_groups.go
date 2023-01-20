package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TicketDefaultGroup struct {
	gorm.Model
	PerformedByID       uuid.UUID `gorm:"type:uuid"`
	TicketType          string    `gorm:"not null"`
	TicketDefaultTypeID uint
	TicketDefaultType   TicketDefaultType
	GroupID             uint `gorm:"not null"`
	Group               Group
	Status              string `gorm:"not null:default:'active'"`
}
