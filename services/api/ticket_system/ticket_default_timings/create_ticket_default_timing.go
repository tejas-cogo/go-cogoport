package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	"fmt"
)

type TicketDefaultTimingService struct {
	TicketDefaultTiming models.TicketDefaultTiming
}

func CreateTicketDefaultTiming(ticket_default_timing models.TicketDefaultTiming) models.TicketDefaultTiming {
	db := config.GetDB()

	fmt.Println(ticket_default_timing)
	// result := map[string]interface{}{}
	db.Create(&ticket_default_timing)
	return ticket_default_timing
}