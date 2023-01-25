package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func DeleteTicket(body models.Ticket) models.Ticket {
	db := config.GetDB()

	var ticket models.Ticket

	db.Model(&ticket).Where("id = ?", body.ID).Update("status", "inactive")

	db.Where("id = ?", body.ID).Delete(&ticket)

	return body
}
