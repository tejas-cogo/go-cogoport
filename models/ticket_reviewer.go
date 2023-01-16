package models

import (
	"gorm.io/gorm"
)

type TicketReviewer struct {
	gorm.Model
	TicketID      uint `gorm:"not null"`
	Ticket        Ticket
	TicketUserID  uint `gorm:"not null"`
	TicketUser    TicketUser
	GroupID       uint `gorm:"not null"`
	Group         Group
	GroupMemberID uint `gorm:"not null"`
	GroupMember   GroupMember
	Status        string `gorm:"default:'active'"`
}
