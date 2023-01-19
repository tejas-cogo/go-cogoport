package models

import (
	"github.com/google/uuid"
	_ "github.com/jinzhu/gorm"
	"gorm.io/gorm"
)

type GroupMember struct {
	gorm.Model
	PerformedByID     uuid.UUID `gorm:"type:uuid"`
	GroupID           uint      `gorm:"not null"`
	Group             Group
	GroupHeadID       uint
	TicketUserID      uint `gorm:"not null"`
	TicketUser        TicketUser
	ActiveTicketCount uint `gorm:"default:0"`
	HierarchyLevel    uint
	Status            string `gorm:"not null:default:'active'"`
}

type CreateGroupMember struct {
	gorm.Model
	PerformedByID  uuid.UUID `gorm:"type:uuid"`
	GroupID        uint
	GroupHeadID    uint
	TicketUserID   []uint `gorm:"type:uint[]"`
	HierarchyLevel uint
}
