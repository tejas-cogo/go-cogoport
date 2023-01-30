package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

type TicketTaskAssigneeService struct {
	TicketTaskAssignee models.TicketTaskAssignee
}

func CreateTicketTaskAssignee(ticket_task_assignee models.TicketTaskAssignee) (string,error,models.TicketTaskAssignee) {
	db := config.GetDB()
	tx := db.Begin()
	var err error

	if err := tx.Create(&ticket_task_assignee).Error; err != nil {
		tx.Rollback()
		return "Error Occurred!", err, ticket_task_assignee
	}

	tx.Commit()
	return "Successfully Created!", err, ticket_task_assignee
}