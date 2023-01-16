package models

import (
	"gorm.io/gorm"
)

type TicketTaskAssignee struct {
	gorm.Model
	TicketID     uint `gorm:"not null"`
	TicketUserID uint `gorm:"not null"`
	Status       string `gorm:"default:'active'"`
}
