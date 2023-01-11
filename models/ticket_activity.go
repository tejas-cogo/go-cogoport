package models

import (
	"gorm.io/gorm"
)

type TicketActivity struct {
	gorm.Model
	TicketID    uint
	TicketUserID uint
	UserType     string
	Description  string
	Data         string
	IsRead       bool
}
