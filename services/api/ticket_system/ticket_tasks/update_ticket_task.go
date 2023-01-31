package ticket_system

import (
	"errors"
	"fmt"

	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func UpdateTicketTask(id uint, body models.TicketTask) (models.TicketTask,error) {
	db := config.GetDB()
	tx := db.Begin()
	var err error

	var ticket_task models.TicketTask
	fmt.Print("Body", body)

	if err := tx.Where("id = ?", id).First(&ticket_task).Error; err != nil {
		tx.Rollback()
		return ticket_task, errors.New("Error Occured!")
	}

	if err := tx.Save(&ticket_task).Error; err != nil {
		tx.Rollback()
		return ticket_task, errors.New("Error Occured!")
	}
	tx.Commit()
	return ticket_task, err
}