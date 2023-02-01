package api

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	"errors"
)

type TicketTaskAssigneeService struct {
	TicketTaskAssignee models.TicketTaskAssignee
}

func CreateTicketTaskAssignee(ticket_task_assignee models.TicketTaskAssignee) (models.TicketTaskAssignee,error) {
	db := config.GetDB()
	tx := db.Begin()
	var err error

	if err := tx.Create(&ticket_task_assignee).Error; err != nil {
		tx.Rollback()
		return ticket_task_assignee, errors.New("Cannot create ticket task assignee!")
	}

	tx.Commit()
	return ticket_task_assignee, err
}