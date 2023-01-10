package models
import (
	"gorm.io/gorm"
 	"github.com/google/uuid"
)
type TicketTask struct {
 	gorm.Model
 	TicketID uint 
 	Title string 
 	CreatedByUserId uuid.UUID
 	Status string 
}