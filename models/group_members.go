package models

import (
	_ "github.com/jinzhu/gorm"
	"gorm.io/gorm"
)

type GroupMember struct {
	gorm.Model
	GroupID           uint
	Group             Group
	TicketUserID      uint
	TicketUser        TicketUser
	ActiveTicketCount uint
	HierarchyLevel    uint
	Status            string
}
