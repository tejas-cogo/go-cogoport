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
	PerformedByID           uuid.UUID
	TicketUser              TicketUser `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Source                  string     `gorm:"not null"`
	Type                    string     `gorm:"not null"`
	Category                string     `gorm:"not null"`
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
	AgentID      string
	AgentRmID    string
	TicketUsers  []uint
	Overdue      int64
	DueToday     int64
	Open         int64
	Escalated    int64
	Rejected     int64
	Closed       int64
	Reassigned   int64
	Unresolved   int64
	ExpiringSoon int64
	HighPriority int64
	StartDate    string
	EndDate      string
}

type TicketGraph struct {
	AgentID     string
	AgentRmID   string
	TodayOpen   TimeDistribution
	TodayClosed TimeDistribution
	WeekOpen    Week
	WeekClosed  Week
	StartDate   time.Time
	EndDate     time.Time
	TodayDate   time.Time
	Sum         int64
}

type TicketEscalatedPayload struct {
	TicketID uint
}

type TimeDistribution struct {
	First  int64
	Second int64
	Third  int64
	Fourth int64
	Fifth  int64
	Sixth  int64
}

type Week struct {
	Monday    int64
	Tuesday   int64
	Wednesday int64
	Thursday  int64
	Friday    int64
	Saturday  int64
	Sunday    int64
}

type TicketExtraFilter struct {
	TicketUserID            uint
	PerformedByID           string
	AgentID                 string
	AgentRmID               string
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
	Status                  string
	TicketCreatedAt         string
	IsExpiringSoon          string
	ExpiryDate              string
	ID                      uint
}
