package models
import (
    "gorm.io/gorm"
)
type TicketTaskAssignee struct {
    gorm.Model
    TicketId uint `gorm:"json:ticket_id"`
    TicketUserId uint `gorm:"json:ticket_user_id"`
    Status string 
}