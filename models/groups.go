package models

import (
	"gorm.io/gorm"
	"github.com/lib/pq"
)

type Group struct {
	gorm.Model
	Name string 
	Tags pq.StringArray `gorm:"type:[]string"`
	Status string 
}