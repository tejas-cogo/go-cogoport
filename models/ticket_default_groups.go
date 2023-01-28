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
	TicketDefaultType   TicketDefaultType `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	GroupID             uint              `gorm:"not null"`
	Group               Group             `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	GroupMemberID       uint
	Status              string `gorm:"not null:default:'active'"`
}
