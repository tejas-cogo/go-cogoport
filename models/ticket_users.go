package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TicketUser struct {
	gorm.Model
	Name         string    `gorm:"not null:json:name:unique"`
	SystemUserID uuid.UUID `gorm:"type:uuid:unique"`
	Email        string    `gorm:"not null:json:email:unique"`
	MobileNumber string    `gorm:"type:varchar(10):unique"`
	RoleID       uint
	Role         Role   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Source       string `gorm:"not null"`
	Type         string `gorm:"not null"`
	Status       string `gorm:"not null:default:'active'"`
}

type TicketUserFilter struct {
	ID              uint
	NotPresentID    uint
	Name            string
	SystemUserID    string
	Email           string
	MobileNumber    string
	RoleID          uint
	Source          string
	Type            string
	Status          string
	RoleUnassigned  bool
	GroupUnassigned bool
}

type TicketUserRole struct {
	ID           []uint
	Name         string
	SystemUserID uuid.UUID
	Email        string
	MobileNumber string
	RoleID       uint
	Source       string
	Type         string
	Status       string
}
