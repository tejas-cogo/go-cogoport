package models
import (
    "gorm.io/gorm"
)
type TicketTaskAssignee struct {
    gorm.Model
    TicketId uint 
    TicketUserId uint
    Status string 
}