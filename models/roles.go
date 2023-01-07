package models

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Name string 
	Level string
	Status string 
}