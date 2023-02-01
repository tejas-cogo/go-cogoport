package api

import (
	"errors"

	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func ListTicketTask() ([]models.TicketTask, error) {
	db := config.GetDB()
	tx := db.Begin()
	var err error

	var ticket_task []models.TicketTask

	if err := tx.Find(&ticket_task).Error; err != nil {
		tx.Rollback()
		return ticket_task, errors.New("Cannot find ticket task!")
	}

	return ticket_task, err
}
