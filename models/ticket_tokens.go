package models

import (
	"gorm.io/gorm"
	"time"
)

type TicketToken struct {
	gorm.Model
	TicketToken        string
	TicketID     uint
	TicketUserID uint
	ExpiryDate   time.Time
	Status       string
}