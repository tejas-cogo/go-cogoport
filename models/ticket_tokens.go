package models

import (
	"time"

	gormjsonb "github.com/dariubs/gorm-jsonb"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type TicketToken struct {
	gorm.Model
	TicketToken  string    `gorm:"not null;unique"`
	TicketID     uint      `gorm:"default;null"`
	TicketUserID uint      `gorm:"not null"`
	ExpiryDate   time.Time `gorm:"not null"`
	Status       string    `gorm:"not null;default:'active'"`
}

type TokenFilter struct {
	TicketToken             string
	Source                  string
	Type                    string
	Category                string
	Subcategory             string
	Description             string
	IsUrgent                bool
	Tags                    pq.StringArray  `gorm:"type:text[]"`
	Data                    gormjsonb.JSONB `gorm:"type:json"`
	NotificationPreferences pq.StringArray  `gorm:"type:text[]"`
}

type TokenActivity struct {
	TicketToken string
	Description string
	Data        gormjsonb.JSONB `gorm:"type:json"`
	Type        string
	Status      string
}
