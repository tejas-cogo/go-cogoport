package models

import (
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

type Group struct {
	gorm.Model
	Name   string
	Tags   pq.StringArray `gorm:"type:text[]"`
	Status string
}
