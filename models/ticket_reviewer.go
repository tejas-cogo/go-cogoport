package models
import (
	"gorm.io/gorm"
)
type TicketReviewer struct {
    gorm.Model
    TicketID uint 
    TicketUserID uint 
    GroupID uint 
    GroupMemberID uint 
    Status string 
}