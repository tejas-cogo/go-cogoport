package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TicketUser struct {
	gorm.Model
	Name         string    `gorm:"not null:json:email"`
	SystemUserID uuid.UUID `gorm:"type:uuid"`
	Email        string    `gorm:"not null:json:email"`
	MobileNumber string    `gorm:"type:varchar(10)"`
	RoleID       uint
	Source       string `gorm:"not null"`
	Type         string `gorm:"not null"`
	Status       string `gorm:"not null:default:'active'"`
}
