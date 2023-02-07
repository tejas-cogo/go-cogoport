package models

import (
	gormjsonb "github.com/dariubs/gorm-jsonb"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type TicketDefaultType struct {
	gorm.Model
	PerformedByID     uuid.UUID       `gorm:"type:uuid" json:"performedByID"`
	TicketType        string          `gorm:"not null;unique"`
	AdditionalOptions gormjsonb.JSONB `gorm:"type:json"`
	ClosureAuthorizer pq.StringArray  `gorm:"type:text[]"`
	Status            string          `gorm:"not null;default:'active'"`
}

type TicketDefault struct {
	ID                    uint
	TicketType            string
	TypeStatus            string
	TicketDefaultTimingID uint
	TimingStatus          string
	ClosureAuthorizer     pq.StringArray `gorm:"type:text[]"`
	ExpiryDuration        string
	Tat                   string
	Conditions            pq.StringArray `gorm:"type:text[]"`
	TicketPriority        string
	AdditionalOptions     gormjsonb.JSONB         `gorm:"type:json"`
	TicketDefaultRole     []TicketTypeDefaultRole `gorm:"foreignKey:TicketDefaultTypeID"`
}

type TicketDefaultFilter struct {
	TicketType string
	QFilter    string
}
