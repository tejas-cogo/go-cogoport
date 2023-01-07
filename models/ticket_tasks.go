package models
import (
	"gorm.io/gorm"
 	"github.com/google/uuid"
)
type TicketTask struct {
 	gorm.Model
 	TicketId Ticket `gorm:"json:ticket_id"`
 	Title string 
 	CreatedByUserId uuid.UUID
 	Status string 
}