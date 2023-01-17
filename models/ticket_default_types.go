package models

import (
	gormjsonb "github.com/dariubs/gorm-jsonb"
	"gorm.io/gorm"
)

type TicketDefaultType struct {
	gorm.Model
	TicketType        string          `gorm:"not null"`
	AdditionalOptions gormjsonb.JSONB `gorm:"type:json"`
	Status            string          `gorm:"not null:default:'active'"`
}
