package models
import (
	"gorm.io/gorm"

)
type TicketDefaultGroup struct {
 	gorm.Model
 	TicketType string 
 	GroupId uint 
 	Status string  
}