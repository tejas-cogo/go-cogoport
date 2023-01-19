package models

import (
	"gorm.io/gorm"
)

type TicketAudit struct {
	gorm.Model
	
	Object   string `gorm:"not null"`
	ObjectId uint `gorm:"not null"`
	Action   string `gorm:"not null"`
	Data     string `gorm:"type:json"`
	Status   string `gorm:"not null:default:'active'"`
}
