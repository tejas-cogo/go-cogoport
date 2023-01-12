package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func ListTicket(filters models.Ticket, tags string) []models.Ticket {
	db := config.GetDB()

	var ticket []models.Ticket

	if filters.Type != "" {
		db = db.Where("type = ?", filters.Type)
	}

	if filters.Priority != "" {
		db = db.Where("priority = ?", filters.Priority)
	}

	if filters.Source != "" {
		db = db.Where("source = ?", filters.Source)
	}

	if tags != "" {
		db = db.Where("? Like ANY(tags)", tags)
	}

	if filters.Status != "" {
		db = db.Where("status = ?", filters.Status)
	} else {
		db = db.Where("status = ?", "active")
	}

	db.Preload("TicketUser").Find(&ticket)

	return ticket
}
