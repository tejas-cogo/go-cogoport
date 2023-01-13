package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type TicketDefaultTiming struct {
	gorm.Model
	TicketType     string
	TicketPriority string
	ExpiryDuration string
	Tat            string
	Conditions     pq.StringArray `gorm:"type:text[]"`
	Status         string
}
