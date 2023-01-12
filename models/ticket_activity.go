package models

import (
	"gorm.io/gorm"
)

type TicketActivity struct {
	gorm.Model
	TicketID   uint
	Ticket Ticket
	TicketUserID uint
	TicketUser TicketUser
	UserType     string
	Description  string
	Data         string
	IsRead       bool
}
