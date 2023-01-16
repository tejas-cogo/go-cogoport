package models

import (
	"time"

	"gorm.io/gorm"
)

type TicketToken struct {
	gorm.Model
	TicketToken  string `gorm:"not null:unique"`
	TicketID     uint   `gorm:"default:null"`
	TicketUserID uint   `gorm:"not null"`
	TicketUser   TicketUser
	ExpiryDate   time.Time `gorm:"not null"`
	Status       string    `gorm:"default:'active'"`
}
