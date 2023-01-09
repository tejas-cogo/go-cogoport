package models

import (
	"gorm.io/gorm"
)

type GoUser struct {
	gorm.Model
	Name string
}
