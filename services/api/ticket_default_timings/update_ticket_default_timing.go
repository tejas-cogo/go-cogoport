package api

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	"errors"
)

func UpdateTicketDefaultTiming(body models.TicketDefaultTiming) (models.TicketDefaultTiming,error) {
	db := config.GetDB()
	tx := db.Begin()
	var err error

	var ticket_default_timing models.TicketDefaultTiming

	if err := tx.Where("id = ?", body.ID).Find(&ticket_default_timing).Error; err != nil {
		tx.Rollback()
		return body, errors.New("Error occured!")
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
		return body, errors.New("Error occured!")
	}

	tx.Commit()

	return ticket_default_timing, err
}
