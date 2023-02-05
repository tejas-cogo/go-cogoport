package api

import (
	"github.com/google/uuid"
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	"gorm.io/gorm"
)

func ListTicketReviewer(filters models.TicketReviewer) ([]models.TicketReviewer, *gorm.DB) {
	db := config.GetDB()

	var ticket_reviewer []models.TicketReviewer

	if filters.TicketID != 0 {
		db = db.Where("ticket_id = ?", filters.TicketID)
	}

	if filters.UserID != uuid.Nil {
		db = db.Where("ticket_user_id = ?", filters.UserID)
	}

	if filters.Status != "" {
		db = db.Where("status = ?", filters.Status)
	}

	db = db.Preload("TicketUser").Find(&ticket_reviewer)

	return ticket_reviewer, db
}
