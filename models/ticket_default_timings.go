package models

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type TicketDefaultTiming struct {
	gorm.Model
	PerformedByID       uuid.UUID `gorm:"type:uuid"`
	TicketType          string    `gorm:"not null"`
	TicketDefaultTypeID uint
	TicketDefaultType   TicketDefaultType `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TicketPriority      string         `gorm:"not null"`
	ExpiryDuration      string         `gorm:"not null"`
	Tat                 string         `gorm:"not null"`
	Conditions          pq.StringArray `gorm:"type:text[]"`
	Status              string         `gorm:"not null:default:'active'"`
}
