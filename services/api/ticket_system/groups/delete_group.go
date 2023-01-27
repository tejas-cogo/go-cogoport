package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func DeleteGroup(id uint) uint{
	db := config.GetDB()

	var group models.Group

	db.Model(&group).Where("id = ?", id).Update("status","inactive")

	db.Where("id = ?", id).Delete(&group)

	return id
}