package models

import (
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Ticket struct {
	gorm.Model
	TicketUserID            uint
	Source                  string
	Type                    string
	Category                string
	Subcategory             string
	Description             string
	Priority                string
	Tags                    pq.StringArray `gorm:"type:text[]"`
	Data                    string
	NotificationPreferences pq.StringArray `gorm:"type:text[]"`
	Tat                     time.Duration  `gorm:"type:string"`
	ExpiryDate              time.Time
	Status                  string
}
