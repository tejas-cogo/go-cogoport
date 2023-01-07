package ticket_system

import (
	"github.com/ChandelShikha/go-cogoport/config"
	"github.com/ChandelShikha/go-cogoport/models"
)

func DeleteTicketDefaultGroup(id uint) uint{
	db := config.GetDB()

	var ticket_default_group models.TicketDefaultGroup

	db.Model(&ticket_default_group).Where("id = ?", id).Update("status","inactive")

	db.Where("id = ?", id).Delete(&ticket_default_group)

	return id
}