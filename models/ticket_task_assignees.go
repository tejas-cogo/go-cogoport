package models

import (
	"gorm.io/gorm"
)

type TicketTaskAssignee struct {
	gorm.Model
	TicketID     uint
	Ticket Ticket
	TicketUserID uint
	TicketUser TicketUser
	Status       string
}
