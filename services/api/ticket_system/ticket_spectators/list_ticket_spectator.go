package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func ListTicketSpectator() []models.TicketSpectator {
	db := config.GetDB()

	var ticket_spectator []models.TicketSpectator

	db.Find(&ticket_spectator)

	return ticket_spectator
}
