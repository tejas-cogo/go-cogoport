package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func ListTicketTask() []models.TicketTask {
	db := config.GetDB()

	var ticket_task []models.TicketTask

	db.Find(&ticket_task)

	return ticket_task
}
