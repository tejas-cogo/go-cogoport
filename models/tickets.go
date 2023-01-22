package models

import (
	"time"

	gormjsonb "github.com/dariubs/gorm-jsonb"
	"github.com/google/uuid"
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
	Priority          string
}

type TicketStat struct {
	AgentID     uuid.UUID `gorm:"type:uuid"`
	AgentRmID   uuid.UUID `gorm:"type:uuid"`
	TicketUsers []uint
	Overdue     int64
	DueToday    int64
	Open        int64
	Escalated   int64
	Rejected    int64
	Closed      int64
	Reassigned  int64
	Unresolved  int64
}
