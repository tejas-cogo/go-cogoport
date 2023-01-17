package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	activities "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_activities"
)

func ReassignTicketSpectator(activity models.Activity, body models.TicketSpectator) models.TicketSpectator {
	db := config.GetDB()
	// result := map[string]interface{}{}

	var ticket_spectator_old []models.TicketSpectator
	var ticket_spectator_active models.TicketSpectator
	var ticket_spectator_new models.TicketSpectator

	ticket_spectator_new.TicketID = body.TicketID

	db.Where("ticket_id = ?", body.TicketID)
	db.Find(&ticket_spectator_old)

	for _, u := range ticket_spectator_old {
		if u == ticket_spectator_new {
			u.Status = "active"
			db.Save(&u)
		}
	}

	db.Where("id = ?", body.ID)
	db.Find(&ticket_spectator_active)

	ticket_spectator_active.Status = "inactive"
	db.Save(&ticket_spectator_active)

	db.Create(&ticket_spectator_new)

	var filters models.Filter

	filters.TicketActivity.TicketID = body.TicketID
	filters.TicketUser.SystemUserID = activity.SystemUserID
	filters.TicketActivity.Type = "Spectator Reassigned"

	activities.CreateTicketActivity(filters)

	return ticket_spectator_new
}
