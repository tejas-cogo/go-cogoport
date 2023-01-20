package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func DeleteTicket(id uint) uint{
	db := config.GetDB()

	var ticket models.Ticket

	db.Model(&ticket).Where("id = ?", id).Update("status","inactive")

	db.Where("id = ?", id).Delete(&ticket)

	return id
}