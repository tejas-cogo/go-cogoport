package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func ListTicketTaskAssignee() []models.TicketTaskAssignee {
	db := config.GetDB()

	var ticket_task_assignee []models.TicketTaskAssignee

	result := map[string]interface{}{}
	db.Find(&ticket_task_assignee).Take(&result)

	return ticket_task_assignee
}
