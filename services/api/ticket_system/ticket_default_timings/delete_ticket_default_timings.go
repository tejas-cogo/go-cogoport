package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func DeleteTicketDefaultTiming(id uint) uint{
	db := config.GetDB()

	var ticket_default_timing models.TicketDefaultTiming

	db.Model(&ticket_default_timing).Where("id = ?", id).Update("status","inactive")

	db.Where("id = ?", id).Delete(&ticket_default_timing)

	return id
}