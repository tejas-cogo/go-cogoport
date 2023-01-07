package ticket_system

import (
	"fmt"
	"github.com/ChandelShikha/go-cogoport/config"
	"github.com/ChandelShikha/go-cogoport/models"
)

func UpdateTicketDefaultTiming(id uint, body models.TicketDefaultTiming) models.TicketDefaultTiming {
	db := config.GetDB()
	var ticket_default_timing models.TicketDefaultTiming
	fmt.Print("Body", body)
	db.Where("id = ?", id).First(&ticket_default_timing)

	// ticket_default_timing.Name = body.Name

	db.Save(&ticket_default_timing)
	return ticket_default_timing
}