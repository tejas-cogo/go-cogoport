package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func UpdateTicket(body models.Ticket) models.Ticket {
	db := config.GetDB()
	var ticket models.Ticket
	db.Where("id = ?", body.ID).First(&ticket)

	db.Save(&ticket)
	return ticket
}
