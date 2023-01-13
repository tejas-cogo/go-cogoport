package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	"gorm.io/gorm"
)

func ListTicketActivity(filters models.TicketActivity) ([]models.TicketActivity,*gorm.DB) {
	db := config.GetDB()

	var ticket_activity []models.TicketActivity

	if filters.TicketID != 0 {
		db = db.Where("ticket_id = ?", filters.TicketID)
	}

	if filters.TicketUserID != 0 {
		db = db.Where("ticket_user_id = ?", filters.TicketUserID)
	}

	if filters.IsRead != false {
		db = db.Where("is_read = ?", filters.IsRead)
	}

	if filters.UserType != "" {
		db = db.Where("user_type = ?", filters.UserType)
	}


	db.Find(&ticket_activity)

	return ticket_activity,db
}
