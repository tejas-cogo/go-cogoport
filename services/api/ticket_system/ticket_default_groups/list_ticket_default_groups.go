package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func ListTicketDefaultGroup(filters models.TicketDefaultGroup) ([]models.TicketDefaultGroup, error) {
	db := config.GetDB()

	var err error

	var ticket_default_groups []models.TicketDefaultGroup

	if filters.TicketDefaultTypeID > 0 {
		db = db.Where("ticket_default_type_id = ?", filters.TicketDefaultTypeID)
	}

	if filters.GroupID != 0 {
		db = db.Where("group_id = ?", filters.GroupID)
	}

	if filters.Status != "" {
		db = db.Where("status = ?", filters.Status)
	}

	if err := db.Order("created_at desc").Find(&ticket_default_groups).Error; err != nil {
		db.Rollback()
		return ticket_default_groups, err
	}



	db = db.Order("created_at desc").Find(&ticket_default_groups)

	return ticket_default_groups, err
}
