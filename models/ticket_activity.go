package models

import (
	gormjsonb "github.com/dariubs/gorm-jsonb"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TicketActivity struct {
	gorm.Model
	TicketID     uint `gorm:"not null"`
	TicketUserID uint
	TicketUser   TicketUser
	UserType     string `gorm:"not null"`
	Description  string
	Type         string          `gorm:"not null"`
	Data         gormjsonb.JSONB `gorm:"type:json"`
	IsRead       bool
	Status       string
}

type Activity struct {
	SystemUserID uuid.UUID
	TicketUserID uint
	Description  string
	Data         gormjsonb.JSONB `gorm:"type:json"`
}
