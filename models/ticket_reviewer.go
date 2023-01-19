package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TicketReviewer struct {
	gorm.Model
	PerformedByID uuid.UUID `gorm:"type:uuid"`
	TicketID      uint      `gorm:"not null"`
	TicketUserID  uint      `gorm:"not null"`
	TicketUser    TicketUser
	GroupID       uint   `gorm:"not null"`
	GroupMemberID uint   `gorm:"not null"`
	Status        string `gorm:"not null:default:'active'"`
}

type ReviewerActivity struct {
	TicketID       uint
	ReviewerUserID uint
	GroupID        uint
	GroupMemberID  uint
	PerformedByID  uuid.UUID `gorm:"type:uuid"`
	Description    string
}
