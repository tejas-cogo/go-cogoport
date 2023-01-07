package ticket_system

import (
	"github.com/ChandelShikha/go-cogoport/config"
	"github.com/ChandelShikha/go-cogoport/models"
)

type TicketSpectatorService struct {
	TicketSpectator models.TicketSpectator
}

func CreateTicketSpectator(ticket_spectator models.TicketSpectator) models.TicketSpectator {
	db := config.GetDB()
	// result := map[string]interface{}{}
	db.Create(&ticket_spectator)
	return ticket_spectator
}