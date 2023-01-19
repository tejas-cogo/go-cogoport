package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TicketTaskAssignee struct {
	gorm.Model
	PerformedByID uuid.UUID `gorm:"type:uuid"`
	TicketID      uint   `gorm:"not null"`
	TicketUserID  uint   `gorm:"not null"`
	Status        string `gorm:"not null:default:'active'"`
}
