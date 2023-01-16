package models

import (
	"gorm.io/gorm"
)

type TicketReviewer struct {
	gorm.Model
	TicketID      uint `gorm:"not null"`
	TicketUserID  uint `gorm:"not null"`
	TicketUser    TicketUser
	GroupID       uint   `gorm:"not null"`
	GroupMemberID uint   `gorm:"not null"`
	Status        string `gorm:"default:'active'"`
}
