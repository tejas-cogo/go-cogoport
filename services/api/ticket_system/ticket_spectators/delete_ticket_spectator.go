package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	activities "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_activities"
)

func DeleteTicketSpectator(id uint) uint {
	db := config.GetDB()

	var ticket_spectator models.TicketSpectator
	var spectator_activity models.SpectatorActivity
	var filters models.Filter

	db.Model(&ticket_spectator).Where("id = ?", id).Update("status", "inactive")

	db.Where("id = ?", id).Delete(&ticket_spectator)

	spectator_activity.TicketID = ticket_spectator.TicketID
	spectator_activity.PerformedByID = ticket_spectator.PerformedByID
	filters.TicketActivity.Type = "Spectator Deleted"
	filters.TicketActivity.Status = "deleted"

	activities.CreateTicketActivity(filters)

	return id
}
