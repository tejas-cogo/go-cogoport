package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Group struct {
	gorm.Model
	Name   string `gorm:"not null"`
	Tags   pq.StringArray `gorm:"type:text[]"`
	Status string `gorm:"default:'active'"`
}
