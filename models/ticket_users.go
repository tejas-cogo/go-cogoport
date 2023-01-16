package models

import (
	"gorm.io/gorm"
)

type TicketUser struct {
	gorm.Model
	Name         string `gorm:"not null:json:email"`
	Email        string `gorm:"not null:json:email"`
	MobileNumber string `gorm:"type:varchar(10)"`
	RoleID       uint 
	Role         Role 
	Source       string `gorm:"not null"`
	Type         string `gorm:"not null"`
	Status       string `gorm:"default:'active'"`
}
