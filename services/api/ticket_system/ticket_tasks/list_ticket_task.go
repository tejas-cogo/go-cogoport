package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func ListTicketTask() []models.TicketTask {
	db := config.GetDB()

	var ticket_task []models.TicketTask

	result := map[string]interface{}{}
	db.Find(&ticket_task).Take(&result)

	return ticket_task
}
