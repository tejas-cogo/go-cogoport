package models

import (
	_"github.com/jinzhu/gorm"
	"gorm.io/gorm"
)

type GroupMember struct {
	gorm.Model
	GroupId Group `gorm:"json:group_id"`
	TicketUserId TicketUser `gorm:"json:ticket_user_id"`
	HierarchyLevel uint
	Status string
}