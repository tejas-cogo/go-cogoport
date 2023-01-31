package ticket_system

import (
	"fmt"
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	"errors"
)

func UpdateTicketTaskAssignee(id uint, body models.TicketTaskAssignee) (models.TicketTaskAssignee,error) {
	db := config.GetDB()
	tx := db.Begin()
	var err error

	var ticket_task_assignee models.TicketTaskAssignee
	fmt.Print("Body", body)

	if err := tx.Where("id = ?", id).First(&ticket_task_assignee).Error; err != nil {
		tx.Rollback()
		return ticket_task_assignee, errors.New("Error Occurred!")
	}

	if err := tx.Save(&ticket_task_assignee).Error; err != nil {
		tx.Rollback()
		return ticket_task_assignee, errors.New("Error Occurred!")
	}

	tx.Commit()
	return ticket_task_assignee, err
}