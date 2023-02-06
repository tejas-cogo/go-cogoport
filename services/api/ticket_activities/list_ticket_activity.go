package api

import (
	"github.com/google/uuid"
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/constants"
	"github.com/tejas-cogo/go-cogoport/models"
	"gorm.io/gorm"
)

func ListTicketActivity(filters models.TicketActivity) ([]models.TicketActivity, *gorm.DB, error) {
	db := config.GetDB()

	var err error

	var ticket_activity []models.TicketActivity

	if filters.TicketID > 0 {
		db = db.Where("ticket_id = ?", filters.TicketID)
	}

	if filters.UserID != uuid.Nil {
		db = db.Where("user_id = ?", filters.UserID)
	}

	if filters.IsRead != false {
		db = db.Where("is_read = ?", filters.IsRead)
	}

	if filters.UserType != "" {

		if filters.UserType == "internal" {
			db = db.Where("type IN ?", constants.AdminActivityView())
		} else if filters.UserType == "client" {
			db = db.Where("type IN ?", constants.ClientActivityView())
		}

	}

	if filters.Status != "" {
		db = db.Where("status = ?", filters.Status)
	}

	// db = db.Order("created_at desc").Preload("TicketUser").Find(&ticket_activity)
	db = db.Order("created_at desc").Find(&ticket_activity)

	return ticket_activity, db, err
}
