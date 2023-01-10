package models

import (
	"time"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type TicketDefaultTiming struct {
	gorm.Model
	TicketType     string
	TicketPriority string
	ExpiryDuration time.Duration
	Tat            time.Duration
	Conditions     pq.StringArray `gorm:"type:text[]"`
	Status         string
}
