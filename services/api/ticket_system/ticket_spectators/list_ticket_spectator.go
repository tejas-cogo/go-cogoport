package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func ListTicketSpectator() []models.TicketSpectator {
	db := config.GetDB()

	var ticket_spectator []models.TicketSpectator

	result := map[string]interface{}{}
	db.Find(&ticket_spectator).Take(&result)

	return ticket_spectator
}
