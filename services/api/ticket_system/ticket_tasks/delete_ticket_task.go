package ticket_system

import (
	"github.com/ChandelShikha/go-cogoport/config"
	"github.com/ChandelShikha/go-cogoport/models"
)

func DeleteTicketTask(id uint) uint{
	db := config.GetDB()

	var ticket_task models.TicketTask

	db.Model(&ticket_task).Where("id = ?", id).Update("status","inactive")

	db.Where("id = ?", id).Delete(&ticket_task)

	return id
}