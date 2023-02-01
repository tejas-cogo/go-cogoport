package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	"gorm.io/gorm"
)

func ListTicketReviewer(filters models.TicketReviewer) ([]models.TicketReviewer, *gorm.DB, error) {
	db := config.GetDB()
	tx := db.Begin()
	var err error

	var ticket_reviewer []models.TicketReviewer

	if filters.TicketID != 0 {
		tx = tx.Where("ticket_id = ?", filters.TicketID)
	}

	if filters.TicketUserID != 0 {
		tx = tx.Where("ticket_user_id = ?", filters.TicketUserID)
	}

	if filters.Status != "" {
		tx = tx.Where("status = ?", filters.Status)
	}

	tx = tx.Preload("TicketUser").Find(&ticket_reviewer)

	tx.Commit()
	return ticket_reviewer, tx, err
}
