package models

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Name   string `gorm:"not null"`
	Level  uint   `gorm:"not null"`
	Status string `gorm:"default:'active'"`
}
