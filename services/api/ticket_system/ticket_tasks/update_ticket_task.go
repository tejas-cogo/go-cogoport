package ticket_system

import (
	"fmt"
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func UpdateTicketTask(id uint, body models.TicketTask) models.TicketTask {
	db := config.GetDB()
	var ticket_task models.TicketTask
	fmt.Print("Body", body)
	db.Where("id = ?", id).First(&ticket_task)

	db.Save(&ticket_task)
	return ticket_task
}