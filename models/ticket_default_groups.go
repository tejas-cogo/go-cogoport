package models
import (
	"gorm.io/gorm"

)
type TicketDefaultGroup struct {
 	gorm.Model
 	TicketType string 
 	GroupID uint
	//Group Group
 	Status string  
}