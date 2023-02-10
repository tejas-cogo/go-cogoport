package models

import (
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type TicketToken struct {
	gorm.Model
	TicketToken  string `gorm:"not null;unique"`
	TicketID     uint
	TicketUserID uint      `gorm:"not null"`
	ExpiryDate   time.Time `gorm:"not null"`
	Status       string    `gorm:"not null; default:'active'"`
}

type TokenFilter struct {
	TicketToken             string
	TicketType              string
	Source                  string
	Type                    string
	Category                string
	Subcategory             string
	Description             string
	IsUrgent                bool
	Tags                    pq.StringArray `gorm:"type:text[]"`
	Data                    Data
	NotificationPreferences pq.StringArray `gorm:"type:text[]"`
	Status                  string
	UserType                string
}

type TokenActivity struct {
	TicketToken string
	Description string
	UserType    string
	Data        Data
	Type        string
	Status      string
}
