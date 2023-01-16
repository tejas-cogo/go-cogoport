package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type TicketDefaultTiming struct {
	gorm.Model
	TicketType     string         `gorm:"not null"`
	TicketPriority string         `gorm:"not null"`
	ExpiryDuration string         `gorm:"not null"`
	Tat            string         `gorm:"not null"`
	Conditions     pq.StringArray `gorm:"type:text[]"`
	Status         string         `gorm:"default:'active'"`
}
