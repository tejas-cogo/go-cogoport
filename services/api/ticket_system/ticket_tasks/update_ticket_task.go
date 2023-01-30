package ticket_system

import (
	"fmt"
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func UpdateTicketTask(id uint, body models.TicketTask) (string,error,models.TicketTask) {
	db := config.GetDB()
	tx := db.Begin()
	var err error

	var ticket_task models.TicketTask
	fmt.Print("Body", body)

	if err := tx.Where("id = ?", id).First(&ticket_task).Error; err != nil {
		tx.Rollback()
		return "Error Occurred!", err, ticket_task
	}

	if err := tx.Save(&ticket_task).Error; err != nil {
		tx.Rollback()
		return "Error Occurred!", err, ticket_task
	}
	return "Successfully Updated!", err, ticket_task
}