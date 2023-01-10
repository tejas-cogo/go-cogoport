package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type TicketUser struct {
	gorm.Model
	Name         string
	Email        string
	MobileNumber string
	RoleID       pq.StringArray `gorm:"type:text[]"`
	Source       string
	Type         string
	Status       string
}
