package models

import (
	"gorm.io/gorm"
)

type TicketSpectator struct {
	gorm.Model
	TicketID     uint `gorm:"not null"`
	Ticket Ticket 
	TicketUserID uint `gorm:"not null"`
	TicketUser TicketUser 
	Status       string `gorm:"default:'active'"`
}
