package models

import (
	"time"

	gormjsonb "github.com/dariubs/gorm-jsonb"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type TicketActivity struct {
	gorm.Model
	TicketID    uint      `gorm:"not null"`
	UserID      uuid.UUID `gorm:"type:uuid"`
	UserType    string    `gorm:"not null"`
	Description string
	Type        string          `gorm:"not null"`
	Data        gormjsonb.JSONB `gorm:"type:json"`
	IsRead      bool
	Status      string
	Ticket      Ticket `gorm:"foreignKey:TicketID"`
}
type TicketActivityData struct {
	TicketID    uint      `gorm:"not null"`
	UserID      uuid.UUID `gorm:"type:uuid"`
	UserType    string    `gorm:"not null"`
	Description string
	Type        string          `gorm:"not null"`
	Data        gormjsonb.JSONB `gorm:"type:json"`
	IsRead      bool
	Status      string
	TicketUser  TicketUser
	Ticket      Ticket `gorm:"foreignKey:TicketID"`
	CreatedAt   time.Time
}

type Activity struct {
	PerformedByID uuid.UUID `gorm:"type:uuid"`
	TicketID      []uint
	UserID        uuid.UUID `gorm:"type:uuid"`
	UserType      string
	Description   string
	Data          gormjsonb.JSONB `gorm:"type:json"`
	Type          string
	Status        string
}

type Communication struct {
	Recipient    string         `json:"recipient"`
	Type         string         `json:"type"`
	Service      string         `json:"service"`
	ServiceID    string         `json:"service_id"`
	TemplateName string         `json:"template_name"`
	Sender       string         `json:"sender"`
	CcEmails     pq.StringArray `gorm:"type:json[]"`
}
