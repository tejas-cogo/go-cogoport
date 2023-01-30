package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func DeleteTicketDefaultGroup(id uint) (string, error, uint){
	db := config.GetDB()
	tx := db.Begin()
	var err error

	var ticket_default_group models.TicketDefaultGroup

	if err := tx.Model(&ticket_default_group).Where("id = ?", id).Update("status","inactive").Error; err != nil {
		tx.Rollback()
		return "Error Occurred!", err ,id
	}

	if err := tx.Where("id = ?", id).Delete(&ticket_default_group).Error; err != nil {
		tx.Rollback()
		return "Error Occurred!", err ,id
	}

	tx.Commit()

	return "Successfully Deleted!", err, id
}