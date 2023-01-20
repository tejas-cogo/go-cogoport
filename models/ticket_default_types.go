package models

import (
	gormjsonb "github.com/dariubs/gorm-jsonb"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type TicketDefaultType struct {
	gorm.Model
	PerformedByID     uuid.UUID       `gorm:"type:uuid"`
	TicketType        string          `gorm:"not null:unique"`
	AdditionalOptions gormjsonb.JSONB `gorm:"type:json"`
	Status            string          `gorm:"not null:default:'active'"`
}

type TicketDefault struct {
	ID                uint
	TicketType        string
	GroupID           uint
	AdditionalOptions gormjsonb.JSONB `gorm:"type:json"`
	TypeStatus        string          `gorm:"not null:default:'active'"`
	TicketPriority    string          `gorm:"not null"`
	ExpiryDuration    string          `gorm:"not null"`
	Tat               string          `gorm:"not null"`
	Conditions        pq.StringArray  `gorm:"type:text[]"`
	TimingStatus      string          `gorm:"not null:default:'active'"`
	GroupName         string          `gorm:"not null:unique"`
	Tags              pq.StringArray  `gorm:"type:text[]"`
	GroupStatus       string          `gorm:"not null:default:'active'"`
}
