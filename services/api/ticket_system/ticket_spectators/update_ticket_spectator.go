package ticket_system

import (
	"fmt"
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func UpdateTicketSpectator(id uint, body models.TicketSpectator) models.TicketSpectator {
	db := config.GetDB()
	var ticket_spectator models.TicketSpectator
	fmt.Print("Body", body)
	db.Where("id = ?", id).First(&ticket_spectator)

	// ticket_spectator.Name = body.Name

	db.Save(&ticket_spectator)
	return ticket_spectator
}