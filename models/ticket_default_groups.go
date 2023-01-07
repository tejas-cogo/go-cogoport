package models
import (
	"gorm.io/gorm"

)
type TicketDefaultGroup struct {
 	gorm.Model
 	TicketType string 
 	GroupId uint `gorm:"json:group_id"`
 	Status string 
}