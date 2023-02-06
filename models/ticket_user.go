package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TicketUser struct {
	gorm.Model
	Name         string    `gorm:"not null:json:name:unique"`
	SystemUserID uuid.UUID `gorm:"type:uuid; unique"`
	Email        string    `gorm:"not null:json:email:unique"`
	MobileNumber string    `gorm:"type:varchar(10); unique"`
	Source       string    `gorm:"not null"`
	Type         string    `gorm:"not null"`
	Status       string    `gorm:"not null; default:'active'"`
}
