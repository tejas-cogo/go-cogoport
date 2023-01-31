package models

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Group struct {
	gorm.Model
	Name          string         `gorm:"not null:unique:json:varchar[]"`
	Tags          pq.StringArray `gorm:"type:text[]"`
	Status        string         `gorm:"not null:default:'active'"`
	PerformedByID uuid.UUID      `gorm:"type:uuid"`
}

type GroupWithMember struct {
	ID            uint
	Name          string
	Tags          pq.StringArray `gorm:"type:text[]"`
	Status        string
	PerformedByID uuid.UUID
	Count         int64
}

type FilterGroup struct {
	ID            uint
	Name          string
	Tags          pq.StringArray
	Status        string
	PerformedByID uuid.UUID
	GroupMemberID uint
}
