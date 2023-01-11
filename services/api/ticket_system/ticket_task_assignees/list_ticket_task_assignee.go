package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func ListTicketTaskAssignee() []models.TicketTaskAssignee {
	db := config.GetDB()

	var ticket_task_assignee []models.TicketTaskAssignee

	db.Find(&ticket_task_assignee)

	return ticket_task_assignee
}
