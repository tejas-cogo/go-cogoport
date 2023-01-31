package ticket_system

import (
	"errors"

	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	helpers "github.com/tejas-cogo/go-cogoport/services/helpers"
)

type TicketDefaultTimingService struct {
	TicketDefaultTiming models.TicketDefaultTiming
}

func CreateTicketDefaultTiming(ticket_default_timing models.TicketDefaultTiming) (models.TicketDefaultTiming, error) {
	db := config.GetDB()
	tx := db.Begin()
	var err error

	stmt := validate(ticket_default_timing)
	if stmt != "validated" {
		return ticket_default_timing, errors.New(stmt)
	}

	ticket_default_timing.Status = "active"

	if err := tx.Create(&ticket_default_timing).Error; err != nil {
		tx.Rollback()
		return ticket_default_timing, errors.New("Error Occurred!")
	}

	tx.Commit()

	return ticket_default_timing, err
}

func validate(ticket_default_timing models.TicketDefaultTiming) string {

	if ticket_default_timing.TicketPriority == "" {
		return ("Ticket Priority Is Required!")
	}

	if ticket_default_timing.ExpiryDuration == "" {
		return ("Expiry Duration Is Required!")
	}

	if ticket_default_timing.Tat == "" {
		return ("Tat Is Required!")
	}

	if ticket_default_timing.TicketDefaultTypeID == 0 {
		return ("Ticket Default Type Is Required!")
	}

	ExpiryDuration := helpers.GetDuration(ticket_default_timing.ExpiryDuration)
	Tat := helpers.GetDuration(ticket_default_timing.Tat)

	if ExpiryDuration <= Tat {
		return ("Expiry Duration should be greater than Tat!")
	}

	return ("validated")
}
