package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	activities "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_activities"
)

type TicketSpectatorService struct {
	TicketSpectator models.TicketSpectator
}

func CreateTicketSpectator(ticket_spectator models.TicketSpectator) models.TicketSpectator {
	db := config.GetDB()
	var spectator_activity models.SpectatorActivity
	var filters models.Filter
	// result := map[string]interface{}{}
	ticket_spectator.Status = "active"

	db.Create(&ticket_spectator)

	spectator_activity.TicketID = ticket_spectator.TicketID
	spectator_activity.PerformedByID = ticket_spectator.PerformedByID
	filters.TicketActivity.Type = "Spectator Assigned"
	filters.TicketActivity.Status = "assigned"

	activities.CreateTicketActivity(filters)
	return ticket_spectator
}
