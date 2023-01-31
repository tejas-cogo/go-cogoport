package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	"gorm.io/gorm"
)

func ListTicketActivity(filters models.TicketActivity) ([]models.TicketActivity, *gorm.DB, error) {
	db := config.GetDB()
	tx := db.Begin()
	var err error

	var ticket_activity []models.TicketActivity

	if filters.TicketID > 0 {
		tx = tx.Where("ticket_id = ?", filters.TicketID)
	}

	if filters.TicketUserID > 0 {
		tx = tx.Where("ticket_user_id = ?", filters.TicketUserID)
	}

	if filters.IsRead != false {
		tx = tx.Where("is_read = ?", filters.IsRead)
	}

	if filters.UserType != "" {
		filters.UserType = "%" + filters.UserType + "%"
		tx = tx.Where("user_type iLike ?", filters.UserType)
	}

	if filters.Status != "" {
		tx = tx.Where("status = ?", filters.Status)
	}

	tx = tx.Order("created_at desc").Preload("TicketUser").Find(&ticket_activity)

	tx.Commit()
	return ticket_activity, tx, err
}
