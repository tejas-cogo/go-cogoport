package models

import (
	"github.com/jinzhu/gorm"
)

type GoUser struct {
	gorm.Model
	id   uint   `gorm:"primayKey", json:"id"`
	Name string `gorm:""json:"name`
}
