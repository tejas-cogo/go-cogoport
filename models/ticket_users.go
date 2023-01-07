package models

import (
	"gorm.io/gorm"
)

type TicketUser struct {
	gorm.Model
	Name string `gorm:"json:name"`
	Email string `gorm:"json:email"`
	MobileNumber string `gorm:"json:mobile_number"`
	RoleIds uint `gorm:"json:role_id"`
	Source string `gorm:"json:source"`
	Type string `gorm:"json:type"`
	Status string `gorm:"json:status"`

}