package models

import (
	_ "github.com/jinzhu/gorm"
	"gorm.io/gorm"
)

type GroupMember struct {
	gorm.Model
	GroupID           uint `gorm:"not null"`
	Group             Group
	TicketUserID      uint `gorm:"not null"`
	TicketUser        TicketUser
	ActiveTicketCount uint `gorm:"default:0"`
	HierarchyLevel    uint
	Status            string `gorm:"default:'active'"`
}
