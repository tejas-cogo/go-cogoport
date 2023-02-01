package api

import (
	"github.com/tejas-cogo/go-cogoport/config"
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

	if filters.TicketUserID > 0 {
		db = db.Where("ticket_user_id = ?", filters.TicketUserID)
	}

	if filters.IsRead != false {
		db = db.Where("is_read = ?", filters.IsRead)
	}

	if filters.UserType != "" {
		filters.UserType = "%" + filters.UserType + "%"
		db = db.Where("user_type iLike ?", filters.UserType)
	}

	if filters.Status != "" {
		db = db.Where("status = ?", filters.Status)
	}

	db = db.Order("created_at desc").Preload("TicketUser").Find(&ticket_activity)


	return ticket_activity, db, err
}
