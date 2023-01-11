package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func DeleteTicketTaskAssignee(id uint) uint{
	db := config.GetDB()

	var ticket_task_assignee models.TicketTaskAssignee

	db.Model(&ticket_task_assignee).Where("id = ?", id).Update("status","inactive")

	db.Where("id = ?", id).Delete(&ticket_task_assignee)

	return id
}