package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type TicketDefaultType struct {
	gorm.Model
	TicketType        string         `gorm:"not null"`
	AdditionalOptions pq.StringArray `gorm:"type:text[]"`
	Status            string         `gorm:"default:'active'"`
}
