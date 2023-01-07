package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

type TicketTaskService struct {
	TicketTask models.TicketTask
}

func CreateTicketTask(ticket_task models.TicketTask) models.TicketTask {
	db := config.GetDB()
	// result := map[string]interface{}{}
	db.Create(&ticket_task)
	return ticket_task
}