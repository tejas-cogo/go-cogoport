package api

import (
	"errors"

	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	validations "github.com/tejas-cogo/go-cogoport/services/validations"
)

type TicketDefaultTimingService struct {
	TicketDefaultTiming models.TicketDefaultTiming
}

func CreateTicketDefaultTiming(ticket_default_timing models.TicketDefaultTiming) (models.TicketDefaultTiming, error) {
	db := config.GetDB()
	tx := db.Begin()
	var err error

	stmt := validations.ValidateTicketDefaultTiming(ticket_default_timing)
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
