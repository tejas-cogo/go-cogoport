package ticket_system

import (
	"fmt"
	"github.com/ChandelShikha/go-cogoport/config"
	"github.com/ChandelShikha/go-cogoport/models"
)

func UpdateTicketTask(id uint, body models.TicketTask) models.TicketTask {
	db := config.GetDB()
	var ticket_task models.TicketTask
	fmt.Print("Body", body)
	db.Where("id = ?", id).First(&ticket_task)

	// ticket_task.Name = body.Name

	db.Save(&ticket_task)
	return ticket_task
}