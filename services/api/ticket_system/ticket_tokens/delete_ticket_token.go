package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func DeleteTicketToken(id uint) uint{
	db := config.GetDB()

	var ticket_token models.TicketToken

	db.Model(&ticket_token).Where("id = ?", id).Update("status","inactive")

	db.Where("id = ?", id).Delete(&ticket_token)

	return id
}