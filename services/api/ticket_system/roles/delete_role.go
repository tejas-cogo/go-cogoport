package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func DeleteRole(id uint) uint{
	db := config.GetDB()

	var role models.Role

	db.Model(&role).Where("id = ?", id).Update("status","inactive")

	db.Where("id = ?", id).Delete(&role)

	return id
}