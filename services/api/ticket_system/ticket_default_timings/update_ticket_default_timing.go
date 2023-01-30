package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func UpdateTicketDefaultTiming(body models.TicketDefaultTiming) (string,error,models.TicketDefaultTiming) {
	db := config.GetDB()
	tx := db.Begin()
	var err error

	var ticket_default_timing models.TicketDefaultTiming

	if err := tx.Where("id = ?", body.ID).Find(&ticket_default_timing).Error; err != nil {
		tx.Rollback()
		return "Error Occurred!", err, body
	}

	if body.TicketDefaultTypeID > 0 {
		ticket_default_timing.TicketDefaultTypeID = body.TicketDefaultTypeID
	}

	if body.TicketPriority != "" {
		ticket_default_timing.TicketPriority = body.TicketPriority
	}
	if body.ExpiryDuration != "" {
		ticket_default_timing.ExpiryDuration = body.ExpiryDuration
	}
	if body.Tat != "" {
		ticket_default_timing.Tat = body.Tat
	}
	if body.Conditions != nil {
		ticket_default_timing.Conditions = body.Conditions
	}
	if body.Status != "" {
		ticket_default_timing.Status = body.Status
	}

	if err := tx.Save(&ticket_default_timing).Error; err != nil {
		tx.Rollback()
		return "Error Occurred!", err, body
	}

	tx.Commit()

	return "Sucessfully Updated!", err, ticket_default_timing
}
