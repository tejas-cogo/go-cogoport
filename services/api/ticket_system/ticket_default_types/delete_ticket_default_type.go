package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func DeleteTicketDefaultType(id uint) uint{
	db := config.GetDB()

	var ticket_default_type models.TicketDefaultType

	db.Model(&ticket_default_type).Where("id = ?", id).Update("status","inactive")

	db.Where("id = ?", id).Delete(&ticket_default_type)

	return id
}