package models

import (
	"time"

	"gorm.io/gorm"
)

type TicketToken struct {
	gorm.Model
	TicketToken  string    `gorm:"not null:unique"`
	TicketID     uint      `gorm:"default:null"`
	TicketUserID uint      `gorm:"not null"`
	ExpiryDate   time.Time `gorm:"not null"`
	Status       string    `gorm:"not null:default:'active'"`
}
