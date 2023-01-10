package models

import (
	"gorm.io/gorm"
)

type TicketActivity struct {
	gorm.Model
	TicketID    uint
	Type         string
	TicketUserID uint
	UserType     string
	Description  string
	Data         string
	IsRead       bool
	Status       string
}
