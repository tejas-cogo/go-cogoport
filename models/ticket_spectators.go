package models

import (
	"gorm.io/gorm"
)

type TicketSpectator struct {
	gorm.Model
	TicketID     uint
	Ticket Ticket
	TicketUserID uint
	TicketUser TicketUser
	Status       string
}
