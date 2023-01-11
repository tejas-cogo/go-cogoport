package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type TicketDefaultTiming struct {
	gorm.Model
	TicketType     string
	TicketPriority string
	ExpiryDuration int
	Tat            int
	Conditions     pq.StringArray `gorm:"type:text[]"`
	Status         string
}
