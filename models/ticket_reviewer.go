package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TicketReviewer struct {
	gorm.Model
	PerformedByID       uuid.UUID `gorm:"type:uuid"`
	TicketID uint      `gorm:"not null"`
	UserID   uuid.UUID `gorm:"not null"`
	RoleID   uuid.UUID `gorm:"not null"`
	// TicketDefaultRoleID uint      `gorm:"not null"`
	Status string `gorm:"not null:default:'active'"`
}

type ReviewerActivity struct {
	TicketID       uint
	ReviewerUserID uuid.UUID
	RoleID         uuid.UUID
	PerformedByID  uuid.UUID `gorm:"type:uuid"`
	Description    string
}
