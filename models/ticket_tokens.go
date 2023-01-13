package models

import (
	"time"

	"gorm.io/gorm"
)

type TicketToken struct {
	gorm.Model
	TicketToken  string
	TicketID    uint 
	//Ticket     Ticket
	TicketUserID uint
	//TicketUser TicketUser
	ExpiryDate   time.Time
	Status       string
}
