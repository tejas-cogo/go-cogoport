

package models
import (
	"gorm.io/gorm"
 	
)
type TicketAudit struct {
 	gorm.Model
	Object string 
	ObjectId uint
	Action string 
	Data string 
	Status string
}