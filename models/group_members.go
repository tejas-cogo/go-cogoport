package models

import (
	_"github.com/jinzhu/gorm"
	"gorm.io/gorm"
)

type GroupMember struct {
	gorm.Model
	GroupId uint `gorm:"json:ticket_id"`
	TicketUserId uint `gorm:"json:ticket_user_id"`
	ActiveTicketCount uint
	HierarchyLevel uint
	Status string
}