package models
import (
	"gorm.io/gorm"
)
type TicketReviewer struct {
    gorm.Model
    TicketId uint 
    TicketUserId uint 
    GroupId uint 
    GroupMemberId uint 
    Status string 
}