package models

import (
	"gorm.io/gorm"
)

type TicketSpectator struct {
	gorm.Model
	TicketID     uint
	TicketUserID uint
	Status       string
}
