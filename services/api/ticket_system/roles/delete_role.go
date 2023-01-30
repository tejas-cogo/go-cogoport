package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func DeleteRole(id uint) uint {
	db := config.GetDB()

	var role models.Role
	var ticket_user models.TicketUser

	db.Model(&role).Where("id = ?", id).Update("status", "inactive")

	db.Model(&ticket_user).Where("role_id = ?", id).Update("role_id", 1)

	db.Where("id = ?", id).Delete(&role)

	return id
}
