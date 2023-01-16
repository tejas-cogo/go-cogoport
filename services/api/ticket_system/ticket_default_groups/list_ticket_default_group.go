package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	"gorm.io/gorm"
	// "fmt"
)

func ListTicketDefaultGroup(filters models.TicketDefaultGroup) ([]models.TicketDefaultGroup, *gorm.DB) {
	db := config.GetDB()

	var ticket_default_groups []models.TicketDefaultGroup

	if filters.TicketType != "" {
		db = db.Where("ticket_type Like ?", filters.TicketType)
	}

	if filters.GroupID != 0 {
		db = db.Where("group_id = ?", filters.GroupID)
	}

	if filters.Status != "" {
		db = db.Where("status = ?", filters.Status)
	} else {
		db = db.Where("status = ?", "active")
	}

	db.Find(&ticket_default_groups)

	return ticket_default_groups, db
}
