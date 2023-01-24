package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	// "fmt"
)

func ListTicketDefaultTiming(filters models.TicketDefaultTiming) ([]models.TicketDefaultTiming, error) {
	db := config.GetDB()
	tx := db.Begin()
	var err error

	var ticket_default_timings []models.TicketDefaultTiming

	if filters.TicketType != "" {
		tx = tx.Where("ticket_type = ?", filters.TicketType)
	
	}

	if filters.TicketPriority != "" {
		tx = tx.Where("ticket_priority = ?", filters.TicketPriority)
	}

	if filters.Status != "" {
		tx = tx.Where("status = ?", filters.Status)
	}

	if err := tx.Order("created_at desc").Find(&ticket_default_timings).Error; err != nil {
		tx.Rollback()
		return ticket_default_timings, err
	}

	tx.Commit()

	return ticket_default_timings, err
}
