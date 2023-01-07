package ticket_system

import (
	"github.com/ChandelShikha/go-cogoport/config"
	"github.com/ChandelShikha/go-cogoport/models"
	// "fmt"
)

func ListTicketDefaultTiming(filters models.TicketDefaultTiming) []models.TicketDefaultTiming {
	db := config.GetDB()

	var ticket_default_timings []models.TicketDefaultTiming

	result := map[string]interface{}{}

	if filters.TicketType != "" {
		db = db.Where("ticket_type = ?", filters.TicketType)
	}

	if filters.TicketPriority != "" {
		db = db.Where("ticket_priority = ?", filters.TicketPriority)
	}

	if filters.Status != "" {
		db = db.Where("status = ?", filters.Status)
	} else {
		db = db.Where("status = ?", "active")
	}

	db.Find(&ticket_default_timings).Take(&result)

	return ticket_default_timings
}
