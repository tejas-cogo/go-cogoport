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
	Status            string          `gorm:"not null;default:'active'"`
}

type TicketDefault struct {
	ID                    uint
	TicketType            string
	TypeStatus            string
	TicketDefaultTimingID uint
	TimingStatus          string
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

type TicketDefaultGroupTypeQuery struct {
	ID                        uint
	TicketDefaultTypeID       uint
	TicketDefaultGroupQueryID uint
	TicketDefaultGroupQuery   []TicketDefaultGroupQuery `gorm:"foreignKey:TicketDefaultGroupID"`
}

type TicketDefaultGroupQuery struct {
	ID                   uint
	TicketDefaultGroupID uint
	GroupLevel           uint
	Status               string
	TicketDefaultTypeID  uint
	GroupQueryID         uint
	GroupQuery           []GroupQuery `gorm:"foreignKey:GroupID"`
	GroupQueryMemberID   uint
	GroupMemberQuery     []GroupMemberQuery `gorm:"foreignKey:GroupMemberID"`
}

type GroupQuery struct {
	ID        uint
	GroupID   uint
	GroupName string
	Count     uint
}

type GroupMemberQuery struct {
	ID               uint
	GroupMemberID    uint
	TicketUserID     uint
	GroupMemberEmail string
	GroupMemberName  string
}
