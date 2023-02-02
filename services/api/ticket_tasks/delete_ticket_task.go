package api

import (
	"errors"

	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func DeleteTicketTask(id uint) (uint, error) {
	db := config.GetDB()
	tx := db.Begin()
	var err error

	var ticket_task models.TicketTask

	if err := tx.Model(&ticket_task).Where("id = ?", id).Update("status", "inactive").Error; err != nil {
		tx.Rollback()
		return id, errors.New("Cannot find ticket task with this id!")
	}

	if err := tx.Where("id = ?", id).Delete(&ticket_task).Error; err != nil {
		tx.Rollback()
		return id, errors.New("Cannot delete ticket task!")
	}

	tx.Commit()
	return id, err
}
