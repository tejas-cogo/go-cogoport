package models
import (
	"gorm.io/gorm"
)
type TicketReviewer struct {
    gorm.Model
    TicketId Ticket `gorm:"json:ticket_id"`
    TicketUserId TicketUser `gorm:"json:ticket_user_id"`
    GroupId Group  `gorm:"json:group_id"`
    GroupMemberId GroupMember `gorm:"json:group_member_id"`
    Status string 
}