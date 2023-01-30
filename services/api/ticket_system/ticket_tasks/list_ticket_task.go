package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func ListTicketTask() (string,error,[]models.TicketTask) {
	db := config.GetDB()
	tx := db.Begin()
	var err error

	var ticket_task []models.TicketTask

	if err := tx.Find(&ticket_task).Error; err != nil {
		tx.Rollback()
		return "Error Occurred!", err, ticket_task
	}

	return "Successfully Listed!", err, ticket_task
}
