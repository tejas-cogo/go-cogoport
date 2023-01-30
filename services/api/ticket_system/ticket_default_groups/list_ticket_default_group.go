package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func ListTicketDefaultGroup(filters models.TicketDefaultGroup) ([]models.TicketDefaultGroup, error) {
	db := config.GetDB()
	tx := db.Begin()
	var err error

	var ticket_default_groups []models.TicketDefaultGroup

	if filters.TicketDefaultTypeID > 0 {
		tx = tx.Where("ticket_default_type_id = ?", filters.TicketDefaultTypeID)
	}

	if filters.GroupID != 0 {
		tx = tx.Where("group_id = ?", filters.GroupID)
	}

	if filters.Status != "" {
		tx = tx.Where("status = ?", filters.Status)
	}

	if err := tx.Order("created_at desc").Find(&ticket_default_groups).Error; err != nil {
		tx.Rollback()
		return ticket_default_groups, err
	}

	tx.Commit()

	db = db.Order("created_at desc").Find(&ticket_default_groups)

	return ticket_default_groups, err
}
