package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func UpdateTicketDefaultTiming(body models.TicketDefaultTiming) models.TicketDefaultTiming {
	db := config.GetDB()
	var ticket_default_timing models.TicketDefaultTiming
	db.Where("id = ?", body.ID).Find(&ticket_default_timing)

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

	db.Save(&ticket_default_timing)
	return ticket_default_timing
}
