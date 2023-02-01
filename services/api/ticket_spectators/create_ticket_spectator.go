package api

import (
	"errors"

	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	activities "github.com/tejas-cogo/go-cogoport/services/api/ticket_activities"
)

type TicketSpectatorService struct {
	TicketSpectator models.TicketSpectator
}

func CreateTicketSpectator(ticket_spectator models.TicketSpectator) (models.TicketSpectator,error) {
	db := config.GetDB()
	tx := db.Begin()
	var err error

	var spectator_activity models.SpectatorActivity
	var filters models.Filter

	ticket_spectator.Status = "active"

	if err := tx.Create(&ticket_spectator).Error; err != nil {
		tx.Rollback()
		return ticket_spectator, errors.New("Error Occurred!")
	}

	spectator_activity.TicketID = ticket_spectator.TicketID
	spectator_activity.PerformedByID = ticket_spectator.PerformedByID
	filters.TicketActivity.Type = "Spectator Assigned"
	filters.TicketActivity.Status = "assigned"

	activities.CreateTicketActivity(filters)

	tx.Commit()
	return ticket_spectator, err
}
