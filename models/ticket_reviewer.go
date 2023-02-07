package models

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type TicketReviewer struct {
	gorm.Model
	PerformedByID      uuid.UUID      `gorm:"type:uuid"`
	TicketID           uint           `gorm:"not null"`
	UserID             uuid.UUID      `gorm:"not null"`
	RoleID             uuid.UUID      `gorm:"not null"`
	ReviewerManagerIDs pq.StringArray `gorm:"type:text[]"`
	Status             string         `gorm:"not null;default:'active'"`
}
type TicketReviewerData struct {
	PerformedByID      uuid.UUID `gorm:"type:uuid"`
	TicketID           uint      `gorm:"not null"`
	UserID             uuid.UUID `gorm:"not null"`
	User               User
	RoleID             uuid.UUID `gorm:"not null"`
	Role               AuthRole
	ReviewerManagerIDs pq.StringArray `gorm:"type:text[]"`
	Status             string         `gorm:"not null;default:'active'"`
}

type ReviewerActivity struct {
	TicketID       uint
	ReviewerUserID uuid.UUID
	RoleID         uuid.UUID
	PerformedByID  uuid.UUID `gorm:"type:uuid"`
	Description    string
}
