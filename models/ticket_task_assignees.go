package models
import (
    "gorm.io/gorm"
)
type TicketTaskAssignee struct {
    gorm.Model
    TicketID uint 
    TicketUserID uint
    Status string 
}