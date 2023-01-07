package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func DeleteTicketActivity(id uint) uint{
	db := config.GetDB()

	var ticket_activity models.TicketActivity

	db.Model(&ticket_activity).Where("id = ?", id).Update("status","inactive")

	db.Where("id = ?", id).Delete(&ticket_activity)

	return id
}