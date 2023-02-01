package api

import (
	"errors"

	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func DeleteTicketTaskAssignee(id uint) (uint, error) {
	db := config.GetDB()
	tx := db.Begin()
	var err error

	var ticket_task_assignee models.TicketTaskAssignee

	if err := tx.Model(&ticket_task_assignee).Where("id = ?", id).Update("status", "inactive").Error; err != nil {
		tx.Rollback()
		return id, errors.New("Error Occurred!")
	}

	if err := tx.Where("id = ?", id).Delete(&ticket_task_assignee).Error; err != nil {
		tx.Rollback()
		return id, errors.New("Error Occurred!")
	}

	tx.Commit()
	return id, err
}
