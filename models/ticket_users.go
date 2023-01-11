package models

import (
	"gorm.io/gorm"
)

type TicketUser struct {
	gorm.Model
	Name         string
	Email        string
	MobileNumber string
	RoleID       uint
	Source       string
	Type         string
	Status       string
}
