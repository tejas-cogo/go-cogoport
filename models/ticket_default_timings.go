package models
import (
	"gorm.io/gorm"
 	"time"
)
type TicketDefaultTiming struct {
 	gorm.Model
	TicketType string 
 	TicketPriority string 
 	ExpiryDuration time.Time
 	Tat time.Time
 	Conditions string 
 	Status string 
}