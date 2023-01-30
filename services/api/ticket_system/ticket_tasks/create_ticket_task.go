package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

type TicketTaskService struct {
	TicketTask models.TicketTask
}

func CreateTicketTask(ticket_task models.TicketTask) (string,error,models.TicketTask) {
	db := config.GetDB()
	tx := db.Begin()
	var err error

	if err := tx.Create(&ticket_task).Error; err != nil {
		tx.Rollback()
		return "Error Occurred!", err, ticket_task
	}

	tx.Commit()
	return "Successfully Created!", err, ticket_task
}