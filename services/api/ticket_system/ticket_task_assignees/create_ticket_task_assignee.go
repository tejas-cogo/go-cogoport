package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

type TicketTaskAssigneeService struct {
	TicketTaskAssignee models.TicketTaskAssignee
}

func CreateTicketTaskAssignee(ticket_task_assignee models.TicketTaskAssignee) models.TicketTaskAssignee {
	db := config.GetDB()
	// result := map[string]interface{}{}
	db.Create(&ticket_task_assignee)
	return ticket_task_assignee
}