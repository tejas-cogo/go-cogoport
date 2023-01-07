package models
import (
	"gorm.io/gorm"
)
type TicketDefaultTypes struct {
 	gorm.Model
 	TicketType string `gorm:"json:ticket_type"`
 	AdditionalOptions string `gorm:"json:additional_optionals"`
 	Status string `gorm:"json:status"`
}