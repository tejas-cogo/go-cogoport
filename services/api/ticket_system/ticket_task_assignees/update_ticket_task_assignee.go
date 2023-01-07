package ticket_system

import (
	"fmt"
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func UpdateTicketTaskAssignee(id uint, body models.TicketTaskAssignee) models.TicketTaskAssignee {
	db := config.GetDB()
	var ticket_task_assignee models.TicketTaskAssignee
	fmt.Print("Body", body)
	db.Where("id = ?", id).First(&ticket_task_assignee)

	// ticket_task_assignee.Name = body.Name

	db.Save(&ticket_task_assignee)
	return ticket_task_assignee
}