package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func DeleteTicketUser(id uint) uint{
	db := config.GetDB()

	var ticket_user models.TicketUser

	db.Model(&ticket_user).Where("id = ?", id).Update("status","inactive")

	db.Where("id = ?", id).Delete(&ticket_user)

	return id
}