package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

type TicketDefaultTimingService struct {
	TicketDefaultTiming models.TicketDefaultTiming
}

func CreateTicketDefaultTiming(ticket_default_timing models.TicketDefaultTiming) (string, error) {
	db := config.GetDB()
	tx := db.Begin()
	var err error

	stmt := validate(ticket_default_timing)
	if stmt != "validated" {
		return stmt, err
	}

	ticket_default_timing.Status = "active"

	if err := tx.Create(&ticket_default_timing).Error; err != nil {
		tx.Rollback()
		return "Error Occurred", err
	}

	tx.Commit()

	return "Successfully Created", err
}

func validate(ticket_default_timing models.TicketDefaultTiming) string {
	if ticket_default_timing.TicketType == "" {
		return ("TicketType Is Required")
	}

	if ticket_default_timing.TicketPriority == "" {
		return ("Ticket Priority Is Required")
	}

	if ticket_default_timing.ExpiryDuration == "" {
		return ("Expiry Duration Is Required")
	}

	if ticket_default_timing.Tat == "" {
		return ("Tat Is Required")
	}

	if ticket_default_timing.TicketDefaultTypeID == 0 {
		return ("Ticket Default Type Is Required")
	}

	return ("validated")
}
