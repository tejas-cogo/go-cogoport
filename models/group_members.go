package models

import (
	_ "github.com/jinzhu/gorm"
	"gorm.io/gorm"
)

type GroupMember struct {
	gorm.Model
	GroupID           uint
	TicketUserID      uint `gorm:"not null"`
	ActiveTicketCount uint `gorm:"default:0"`
	HierarchyLevel    uint
	Status            string `gorm:"default:'active'"`
}
