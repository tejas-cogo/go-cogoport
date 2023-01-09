package models
import (
    "gorm.io/gorm"

)
type TicketSpectator struct {
    gorm.Model
    TicketId uint 
    TicketUserId uint 
    Status string 
}