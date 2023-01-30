package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func DeleteTicket(body models.Ticket) (string,error,models.Ticket) {
	db := config.GetDB()
	tx := db.Begin()
	var err error

	var ticket models.Ticket

	if err := tx.Model(&ticket).Where("id = ?", body.ID).Update("status", "inactive").Error; err != nil {
		tx.Rollback()
		return "Error Occurred!", err, body
	}

	if err := tx.Where("id = ?", body.ID).Delete(&ticket).Error; err != nil {
		tx.Rollback()
		return "Error Occurred!", err, body
	}

	tx.Commit()
	return "Successfully Deleted!", err, body
}
