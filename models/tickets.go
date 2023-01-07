package models
import (
	"gorm.io/gorm"
)
type Ticket struct {
 	gorm.Model
 	TicketUserId TicketUser `gorm:"json:ticket_user_id"`
 	Source string `gorm:"json:source"`
 	Type string `gorm:"json:type"`
 	Category string `gorm:"json:category"`
 	Subcategory string `gorm:"json:subcategory"`
 	Description string `gorm:"json:description"`
 	Priority string `gorm:"json:priority"`
 	Tags string `gorm:"json:tags"`
 	Data string `gorm:"json:data"`
 	NotificationPreferences string `gorm:"json:notification_preferences"`
 	Tat string `gorm:"json:tat"`
 	ExpiryDate string `gorm:"json:expiry_date"`
 	Status string `gorm:"json:status"`
}