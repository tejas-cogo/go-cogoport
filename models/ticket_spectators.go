package models

import (
	"gorm.io/gorm"
)

type TicketSpectator struct {
	gorm.Model
	TicketID     uint `gorm:"not null"`
	TicketUserID uint `gorm:"not null"`
	Status       string `gorm:"default:'active'"`
}

type Filter struct {
	gorm.Model
	Ticket              Ticket
	TicketUser          TicketUser
	Group               Group
	GroupMember         GroupMember
	Role                Role
	TicketActivity      TicketActivity
	TicketAudit         TicketAudit
	TicketDefaultGroup  TicketDefaultGroup
	TicketDefaultTiming TicketDefaultTiming
	TicketDefaultType   TicketDefaultType
	TicketReviewer      TicketReviewer
	TicketSpectator     TicketSpectator
	TicketTask          TicketTask
	TicketToken         TicketToken
}
