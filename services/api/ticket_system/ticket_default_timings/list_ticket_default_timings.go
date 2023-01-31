package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	"gorm.io/gorm"
)

func ListTicketDefaultTiming(filters models.TicketDefaultTiming) ([]models.TicketDefaultTiming, *gorm.DB, error) {
	db := config.GetDB()
	tx := db.Begin()
	var err error

	var ticket_default_timings []models.TicketDefaultTiming

	if filters.TicketDefaultTypeID > 0 {
		tx = tx.Where("ticket_default_type_id = ?", filters.TicketDefaultTypeID)

	}

	if filters.TicketPriority != "" {
		tx = tx.Where("ticket_priority = ?", filters.TicketPriority)
	}

	if filters.Status != "" {
		tx = tx.Where("status = ?", filters.Status)
	}

	if err := tx.Order("created_at desc").Find(&ticket_default_timings).Error; err != nil {
		tx.Rollback()
		return ticket_default_timings, tx, err
	}

	tx.Commit()

	return ticket_default_timings, tx, err
}
