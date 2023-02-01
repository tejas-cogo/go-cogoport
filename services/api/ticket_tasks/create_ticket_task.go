package api

import (
	"errors"

	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

type TicketTaskService struct {
	TicketTask models.TicketTask
}

func CreateTicketTask(ticket_task models.TicketTask) (models.TicketTask, error) {
	db := config.GetDB()
	tx := db.Begin()
	var err error

	if err := tx.Create(&ticket_task).Error; err != nil {
		tx.Rollback()
		return ticket_task, errors.New("Error Occurred!")
	}

	tx.Commit()
	return ticket_task, err
}
