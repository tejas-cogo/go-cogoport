package models

import (
	"time"

	gormjsonb "github.com/dariubs/gorm-jsonb"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Ticket struct {
	gorm.Model
	TicketUserID            uint `gorm:"not null"`
	TicketUser              TicketUser
	Source                  string `gorm:"not null"`
	Type                    string `gorm:"not null"`
	Category                string `gorm:"not null"`
	Subcategory             string
	Description             string
	Priority                string          `gorm:"not null:default:'low'"`
	Tags                    pq.StringArray  `gorm:"type:text[]"`
	Data                    gormjsonb.JSONB `gorm:"type:json"`
	NotificationPreferences pq.StringArray  `gorm:"type:text[]"`
	Tat                     string          `gorm:"not null"`
	ExpiryDate              time.Time       `gorm:"not null"`
	Status                  string          `gorm:"not null:default:'active'"`
}

type TicketDetail struct {
	TicketReviewerID  uint
	TicketReviewer    TicketReviewer
	TicketSpectatorID uint
	TicketSpectator   TicketSpectator
	TicketActivityID  uint
	TicketActivity    []TicketActivity
	TicketID          uint
	Ticket            Ticket
}

type TicketStat struct {
	PerformedByID uint
	Overdue       uint
	DueToday      uint
	Open          uint
	Escalated     uint
	Rejected      uint
	Closed        uint
	Reassigned    uint
}
