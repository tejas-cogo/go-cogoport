package models

import (
	"gorm.io/gorm"
)

type TicketSpectator struct {
	gorm.Model
	TicketID     uint `gorm:"not null"`
	TicketUserID uint `gorm:"not null"`
	TicketUser   TicketUser
	Status       string `gorm:"not null:default:'active'"`
}

type SpectatorActivity struct {
	Activity       Activity
	TicketSpectator TicketSpectator
}

