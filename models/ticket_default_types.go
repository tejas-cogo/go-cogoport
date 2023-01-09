package models
import (
	"gorm.io/gorm"
	"github.com/lib/pq"
)
type TicketDefaultType struct {
 	gorm.Model
 	TicketType string 
 	AdditionalOptions pq.StringArray `gorm:"type:[]string"`
 	Status string 
}