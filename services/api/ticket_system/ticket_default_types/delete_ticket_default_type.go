package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func DeleteTicketDefaultType(id uint) uint {
	db := config.GetDB()

	var ticket_default_type models.TicketDefaultType
	var ticket_default_group models.TicketDefaultGroup
	var ticket_default_timing models.TicketDefaultTiming

	db.Model(&ticket_default_type).Where("id = ?", id).Update("status", "inactive")

	db.Model(&ticket_default_group).Where("ticket_default_type_id = ?", id).Update("status", "inactive")
	db.Where("ticket_default_type_id = ?", id).Delete(&ticket_default_group)

	db.Model(&ticket_default_timing).Where("ticket_default_type_id = ?", id).Update("status", "inactive")
	db.Where("ticket_default_type_id = ?", id).Delete(&ticket_default_timing)

	db.Where("id = ?", id).Delete(&ticket_default_type)

	return id
}
