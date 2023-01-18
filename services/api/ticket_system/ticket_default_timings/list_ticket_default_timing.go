package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	"gorm.io/gorm"
	// "fmt"
)

func ListTicketDefaultTiming(filters models.TicketDefaultTiming) ([]models.TicketDefaultTiming, *gorm.DB) {
	db := config.GetDB()

	var ticket_default_timings []models.TicketDefaultTiming

	if filters.TicketType != "" {
		filters.TicketType = "%" + filters.TicketType + "%"
		db = db.Where("ticket_type Like ?", filters.TicketType)
	}

	if filters.TicketPriority != "" {
		db = db.Where("ticket_priority = ?", filters.TicketPriority)
	}

	if filters.Status != "" {
		db = db.Where("status = ?", filters.Status)
	}

	db = db.Order("created_at desc").Find(&ticket_default_timings)

	return ticket_default_timings, db
}
