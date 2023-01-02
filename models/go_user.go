package models

import (
	"github.com/jinzhu/gorm"
)

type GoUser struct {
	gorm.Model
	ID   uint   `gorm:"primayKey", json:"id"`
	Name string `gorm:""json:"name`
}
