package ticket_system

import (
	"github.com/ChandelShikha/go-cogoport/config"
	"github.com/ChandelShikha/go-cogoport/models"
	// "time"
)

func DeleteGroup(id uint) uint{
	db := config.GetDB()

	var group models.Group

	// db.Where("id = ?", id).First(&group)
	// group.Status = "inactive"
	// db.Save(&group)
	db.Model(&group).Where("id = ?", id).Update("status","inactive")

	db.Where("id = ?", id).Delete(&group)

	return id
}