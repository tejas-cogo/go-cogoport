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
	ID                       uint
	TicketType               string
	AdditionalOptions        gormjsonb.JSONB `gorm:"type:json"`
	TypeStatus               string
	TicketDefaultTimingID    uint
	TicketPriority           string
	ExpiryDuration           string
	Tat                      string
	Conditions               pq.StringArray `gorm:"type:text[]"`
	TimingStatus             string
	TicketDefaultGroupID     uint
	GroupID                  uint
	GroupName                string
	GroupMemberName          string
	GroupMemberID            uint
	Tags                     pq.StringArray `gorm:"type:text[]"`
	MemberCount              int
	TicketDefaultGroupStatus string
}
