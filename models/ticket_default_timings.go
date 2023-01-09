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
	ExpiryDuration time.Time
	Tat            time.Time
	Conditions     pq.StringArray `gorm:"type:[]string"`
	Status         string
}
