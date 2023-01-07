package models
import (
    "gorm.io/gorm"

)
type TicketSpectator struct {
    gorm.Model
    TicketId Ticket `gorm:"json:ticket_id"`
    TicketUserId TicketUser `gorm:"json:ticket_user_id"`
    Status string 
}