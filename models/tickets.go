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
	TicketUserID            uint
	UserID                  uuid.UUID `gorm:"type:uuid"`
	UserType                string
	Source                  string `gorm:"not null"`
	Type                    string `gorm:"not null"`
	Category                string
	Subcategory             string
	Description             string
	Priority                string          `gorm:"not null;default:'low'"`
	Tags                    pq.StringArray  `gorm:"type:text[]"`
	Data                    gormjsonb.JSONB `gorm:"type:json"`
	NotificationPreferences pq.StringArray  `gorm:"type:text[]"`
	Tat                     time.Time       `gorm:"not null"`
	ExpiryDate              time.Time       `gorm:"not null"`
	IsUrgent                bool
	Status                  string `gorm:"not null;default:'active'"`
}

type TicketData struct {
	ID                      uint
	TicketUserID            uint
	UserID                  uuid.UUID `gorm:"type:uuid"`
	User                    User
	UserType                string
	Source                  string `gorm:"not null"`
	Type                    string `gorm:"not null"`
	Category                string
	Subcategory             string
	Description             string
	Priority                string          `gorm:"not null;default:'low'"`
	Tags                    pq.StringArray  `gorm:"type:text[]"`
	Data                    gormjsonb.JSONB `gorm:"type:json"`
	NotificationPreferences pq.StringArray  `gorm:"type:text[]"`
	Tat                     time.Time       `gorm:"not null"`
	ExpiryDate              time.Time       `gorm:"not null"`
	IsUrgent                bool
	Status                  string `gorm:"not null;default:'active'"`
}

type TicketDetail struct {
	TicketReviewerID  uint
	TicketReviewer    TicketReviewerData
	TicketSpectatorID uint
	TicketSpectator   TicketSpectator
	TicketID          uint
	Ticket            Ticket
	TicketUser        TicketUser
}

type TicketStat struct {
	AgentID         string
	AgentRmID       string
	TicketUsers     []uint
	Overdue         int64
	DueToday        int64
	Open            int64
	Escalated       int64
	Rejected        int64
	Closed          int64
	Reassigned      int64
	Unresolved      int64
	ExpiringSoon    int64
	HighPriority    int64
	StartDate       string
	EndDate         string
	UserID          string
	QFilter         string
	ExpiryDate      string
	TicketCreatedAt string
	Tags            pq.StringArray `gorm:"type:text[]"`
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
	QFilter                 string
	PerformedByID           string
	MyTicket                string
	AgentID                 string
	AgentRmID               string
	UserID                  string
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
	ExpiringSoon            string
	ExpiryDate              string
	ID                      uint
}
