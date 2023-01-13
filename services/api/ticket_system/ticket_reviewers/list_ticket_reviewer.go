package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	"gorm.io/gorm"
)

func ListTicketReviewer(filters models.TicketReviewer) ([]models.TicketReviewer,*gorm.DB) {
	db := config.GetDB()

	var ticket_reviewer []models.TicketReviewer

	if filters.TicketID != 0 {
		db = db.Where("ticket_id = ?", filters.TicketID)
	}

	if filters.TicketUserID != 0 {
		db = db.Where("ticket_user_id = ?", filters.TicketUserID)
	}

	db.Find(&ticket_reviewer)

	return ticket_reviewer,db
}
