package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func DeleteTicketSpectator(id uint) uint{
	db := config.GetDB()

	var ticket_spectator models.TicketSpectator

	db.Model(&ticket_spectator).Where("id = ?", id).Update("status","inactive")

	db.Where("id = ?", id).Delete(&ticket_spectator)

	return id
}