package models

import (
	_"github.com/jinzhu/gorm"
	"gorm.io/gorm"
)

type GroupMember struct {
	gorm.Model
	GroupID uint 
	TicketUserID uint
	ActiveTicketCount uint
	HierarchyLevel uint
	Status string
}