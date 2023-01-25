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
	Group             Group `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	GroupHeadID       uint `gorm:"not null"`
	TicketUserID      uint `gorm:"not null"`
	TicketUser        TicketUser `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ActiveTicketCount uint   `gorm:"default:0"`
	HierarchyLevel    uint   `gorm:"not null"`
	Status            string `gorm:"not null:default:'active'"`
}

type CreateGroupMember struct {
	gorm.Model
	PerformedByID  uuid.UUID `gorm:"type:uuid"`
	GroupID        uint      `gorm:"not null"`
	GroupHeadID    uint      `gorm:"not null"`
	TicketUserID   []uint    `gorm:"type:uint[]"`
	HierarchyLevel uint      `gorm:"not null"`
}

type FilterGroupMember struct {
	ID                uint
	PerformedByID     string
	GroupID           uint
	Group             Group 
	GroupHeadID       uint
	TicketUserID      uint
	TicketUser        TicketUser 
	ActiveTicketCount uint
	HierarchyLevel    uint
	Status            string
	NotGroupID        uint
}
