package models

import (
	"gorm.io/gorm"
)

type TicketActivity struct {
	gorm.Model
	TicketId     uint 
	Type         string
	TicketUserId uint 
	UserType     string
	Description  string
	Data         string
	IsRead       bool
	Status       string
}
