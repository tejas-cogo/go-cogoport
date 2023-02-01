package api

import (
	"errors"

	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func UpdateTicketTask(id uint, body models.TicketTask) (models.TicketTask, error) {
	db := config.GetDB()
	tx := db.Begin()
	var err error

	var ticket_task models.TicketTask

	if err := tx.Where("id = ?", id).First(&ticket_task).Error; err != nil {
		tx.Rollback()
		return ticket_task, errors.New("Cannot find ticket task with this id!")
	}

	if err := tx.Save(&ticket_task).Error; err != nil {
		tx.Rollback()
		return ticket_task, errors.New("Cannot save ticket task!")
	}
	tx.Commit()
	return ticket_task, err
}
