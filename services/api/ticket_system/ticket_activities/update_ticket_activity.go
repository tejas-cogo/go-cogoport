package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func UpdateTicketActivity(body models.TicketActivity) (string,error,models.TicketActivity) {
	db := config.GetDB()
	tx := db.Begin()
	var err error

	var ticket_activity models.TicketActivity
	
	if err := tx.Where("id = ?", body.ID).Find(&ticket_activity).Error; err != nil {
		tx.Rollback()
		return "Error Occurred!", err, body
	}

	if err := tx.Save(&ticket_activity).Error; err != nil {
		tx.Rollback()
		return "Error Occurred!", err, body
	}

	tx.Commit()
	
	return "Sucessfully Updated!", err, ticket_activity
}
