package ticket_system

import (
	"fmt"

	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	activities "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_activities"
	// activities "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_activities"
)

func ReassignTicketSpectator(body models.SpectatorActivity) string {
	db := config.GetDB()
	// result := map[string]interface{}{}

	var ticket_spectator_old models.TicketSpectator
	var ticket_spectator_active models.TicketSpectator

	db.Where("ticket_id = ? AND status = 'active'", body.TicketID).Find(&ticket_spectator_active)

	ticket_spectator_active.Status = "inactive"
	db.Save(&ticket_spectator_active)

	fmt.Println("edcs", body, "rfd")

	db.Where("ticket_id = ? AND ticket_user_id = ?", body.TicketID, body.SpectatorUserID).Find(&ticket_spectator_old)

	if ticket_spectator_old.ID != 0 {
		ticket_spectator_old.Status = "active"
		db.Save(&ticket_spectator_old)
	} else {
		var ticket_spectator models.TicketSpectator
		ticket_spectator.TicketID = body.TicketID
		ticket_spectator.TicketUserID = body.SpectatorUserID
		db.Create(&ticket_spectator)
	}

	var filters models.Filter

	filters.TicketActivity.TicketID = body.TicketID
	filters.TicketUser.SystemUserID = body.PerformedByID
	filters.TicketActivity.Type = "Spectator Reassigned"
	filters.TicketActivity.Description = body.Description
	filters.TicketActivity.Status = "reassigned"
	activities.CreateTicketActivity(filters)

	return "Reassigned Successfully"
}
