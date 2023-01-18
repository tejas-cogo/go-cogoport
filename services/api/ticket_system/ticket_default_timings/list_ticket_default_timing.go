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
		db.Where("ticket_type = ?", filters.TicketType)
	}

	if filters.TicketPriority != "" {
		db.Where("ticket_priority = ?", filters.TicketPriority)
	}

	db = db.Find(&ticket_default_timings)

	return ticket_default_timings, db
}
