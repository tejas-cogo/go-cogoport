package models
import (
	"gorm.io/gorm"
 	"github.com/google/uuid"
)
type TicketTask struct {
 	gorm.Model
 	TicketId uint 
 	Title string 
 	CreatedByUserId uuid.UUID
 	Status string 
}