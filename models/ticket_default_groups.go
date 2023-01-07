package models
import (
	"gorm.io/gorm"

)
type TicketDefaultGroup struct {
 	gorm.Model
 	TicketType string 
 	GroupId Group `gorm:"json:group_id"`
 	Status string 
}