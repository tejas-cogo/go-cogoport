package models

import (
	"gorm.io/gorm"
)

type TicketDefaultGroup struct {
	gorm.Model
	TicketType string `gorm:"not null"`
	GroupID    uint   `gorm:"not null"`
	Group      Group
	Status     string `gorm:"default:'active'"`
}
