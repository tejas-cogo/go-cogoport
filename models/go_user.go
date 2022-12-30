package models

import (
	"go-cogoport/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.db

type GoUser struct {
	gorm.model
	id   uint   `gorm:"primayKey"`
	Name string `gorm:""json:"name"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&GoUser{})
}

func GetAllUsers() []GoUser{
	var Users []GoUser
	db.First(&go_user)
	return Users
}
