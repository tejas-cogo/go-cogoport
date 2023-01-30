package models

import (
	"time"

	gormjsonb "github.com/dariubs/gorm-jsonb"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type TicketToken struct {
	gorm.Model
	TicketToken             string     `gorm:"not null:unique"`
	TicketID                uint       `gorm:"default:null"`
	TicketUserID            uint       `gorm:"not null"`
	ExpiryDate              time.Time  `gorm:"not null"`
	Status                  string     `gorm:"not null:default:'active'"`
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
}
