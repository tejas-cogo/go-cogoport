package models

import (
	"time"

	gormjsonb "github.com/dariubs/gorm-jsonb"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Ticket struct {
	gorm.Model
	TicketUserID            uint
	TicketUser              TicketUser
	Source                  string
	Type                    string
	Category                string
	Subcategory             string
	Description             string
	Priority                string
	Tags                    pq.StringArray  `gorm:"type:text[]"`
	Data                    gormjsonb.JSONB `gorm:"type:json"`
	NotificationPreferences pq.StringArray  `gorm:"type:text[]"`
	Tat                     string
	ExpiryDate              time.Time
	Status                  string
}
