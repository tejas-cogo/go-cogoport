package models

import (
	"gorm.io/gorm"
)

type TicketReviewer struct {
	gorm.Model
	TicketID      uint
	Ticket        Ticket
	TicketUserID  uint
	TicketUser    TicketUser
	GroupID       uint
	Group         Group
	GroupMemberID uint
	GroupMember   GroupMember
	Status        string
}
