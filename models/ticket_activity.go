package models

import (
	gormjsonb "github.com/dariubs/gorm-jsonb"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type TicketActivity struct {
	gorm.Model
	TicketID     uint `gorm:"not null"`
	TicketUserID uint
	TicketUser   TicketUser `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserType     string     `gorm:"not null"`
	Description  string
	Type         string          `gorm:"not null"`
	Data         gormjsonb.JSONB `gorm:"type:json"`
	IsRead       bool
	Status       string
}

type Activity struct {
	PerformedByID uuid.UUID `gorm:"type:uuid"`
	TicketID      []uint
	TicketUserID  uint
	Description   string
	Data          gormjsonb.JSONB `gorm:"type:json"`
	Type          string
	Status        string
}

type Post struct {
	Recipient    string         `json:"recipient"`
	Type         string         `json:"type"`
	Service      string         `json:"service"`
	ServiceID    string         `json:"service_id"`
	TemplateName string         `json:"template_name"`
	Sender       string         `json:"sender"`
	CcEmails     pq.StringArray `gorm:"type:json[]"`
}
