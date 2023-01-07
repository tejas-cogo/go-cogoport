package models

import (
	"gorm.io/gorm"
)

type TicketActivity struct {
	gorm.Model
	TicketId     Ticket `gorm:"json:ticket_id"`
	Type         string
	TicketUserId TicketUser `gorm:"json:ticket_user_id"`
	UserType     string
	Description  string
	Data         string
	IsRead       bool
	Status       string
}
