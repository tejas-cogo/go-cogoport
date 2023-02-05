package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TicketReviewer struct {
	gorm.Model
	PerformedByID       uuid.UUID `gorm:"type:uuid"`
	TicketID            uint      `gorm:"not null"`
	UserID              string    `gorm:"not null"`
	RoleID              string    `gorm:"not null"`
	TicketDefaultRoleID uint      `gorm:"not null"`
	Status              string    `gorm:"not null:default:'active'"`
}

type ReviewerActivity struct {
	TicketID       uint
	ReviewerUserID uint
	GroupID        uint
	GroupMemberID  uint
	PerformedByID  uuid.UUID `gorm:"type:uuid"`
	Description    string
}
