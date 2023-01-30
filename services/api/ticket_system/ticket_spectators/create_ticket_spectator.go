package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	activities "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_activities"
)

type TicketSpectatorService struct {
	TicketSpectator models.TicketSpectator
}

func CreateTicketSpectator(ticket_spectator models.TicketSpectator) (string,error,models.TicketSpectator) {
	db := config.GetDB()
	tx := db.Begin()
	var err error

	var spectator_activity models.SpectatorActivity
	var filters models.Filter

	ticket_spectator.Status = "active"

	if err := tx.Create(&ticket_spectator).Error; err != nil {
		tx.Rollback()
		return "Error Occurred!", err, ticket_spectator
	}

	spectator_activity.TicketID = ticket_spectator.TicketID
	spectator_activity.PerformedByID = ticket_spectator.PerformedByID
	filters.TicketActivity.Type = "Spectator Assigned"
	filters.TicketActivity.Status = "assigned"

	activities.CreateTicketActivity(filters)

	tx.Commit()
	return "Successfully Created!", err, ticket_spectator
}
