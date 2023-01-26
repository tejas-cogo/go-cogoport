package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func DeleteTicketUser(id uint) uint {
	db := config.GetDB()

	var ticket_user models.TicketUser

	db.Model(&ticket_user).Where("id = ?", id).Update("role_id", 1)

	return id
}
