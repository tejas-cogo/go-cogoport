package models

import (
	gormjsonb "github.com/dariubs/gorm-jsonb"
	"gorm.io/gorm"
)

type TicketActivity struct {
	gorm.Model
	TicketID     uint `gorm:"not null"`
	TicketUserID uint `gorm:"not null"`
	UserType     string `gorm:"not null"`
	Description  string
	Type         string          `gorm:"not null"`
	Data         gormjsonb.JSONB `gorm:"type:json"`
	IsRead       bool            `gorm:"not null"`
}
