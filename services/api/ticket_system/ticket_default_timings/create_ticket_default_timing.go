package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

type TicketDefaultTimingService struct {
	TicketDefaultTiming models.TicketDefaultTiming
}

func CreateTicketDefaultTiming(ticket_default_timing models.TicketDefaultTiming) models.TicketDefaultTiming {
	db := config.GetDB()
	ticket_default_timing.Status = "active"
	// result := map[string]interface{}{}
	db.Create(&ticket_default_timing)
	return ticket_default_timing
}
