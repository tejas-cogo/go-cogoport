package api

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	"gorm.io/gorm"
)

func ListTicketDefaultTiming(filters models.TicketDefaultTiming) ([]models.TicketDefaultTiming, *gorm.DB, error) {
	db := config.GetDB()

	var err error

	var ticket_default_timings []models.TicketDefaultTiming

	if filters.TicketDefaultTypeID > 0 {
		db = db.Where("ticket_default_type_id = ?", filters.TicketDefaultTypeID)

	}

	if filters.TicketPriority != "" {
		db = db.Where("ticket_priority = ?", filters.TicketPriority)
	}

	if filters.Status != "" {
		db = db.Where("status = ?", filters.Status)
	}

	db = db.Order("created_at desc").Find(&ticket_default_timings)

	return ticket_default_timings, db, err
}
