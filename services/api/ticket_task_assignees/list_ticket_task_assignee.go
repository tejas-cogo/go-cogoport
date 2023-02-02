package api

import (
	"errors"

	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func ListTicketTaskAssignee() ([]models.TicketTaskAssignee, error) {
	db := config.GetDB()
	tx := db.Begin()
	var err error

	var ticket_task_assignee []models.TicketTaskAssignee

	if err := tx.Find(&ticket_task_assignee).Error; err != nil {
		tx.Rollback()
		return ticket_task_assignee, errors.New(err.Error())
	}

	tx.Commit()
	return ticket_task_assignee, err
}
