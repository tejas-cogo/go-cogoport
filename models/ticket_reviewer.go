package models
import (
	"gorm.io/gorm"
)
type TicketReviewer struct {
    gorm.Model
    TicketId uint `gorm:"json:ticket_id"`
    TicketUserId uint `gorm:"json:ticket_user_id"`
    GroupId Group  `gorm:"json:group_id"`
    GroupMemberId uint `gorm:"json:group_member_id"`
    Status string 
}