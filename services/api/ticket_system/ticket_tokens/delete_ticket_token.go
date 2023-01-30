package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func DeleteTicketToken(id uint) (string,error,uint){
	db := config.GetDB()
	tx := db.Begin()
	var err error

	var ticket_token models.TicketToken

	if err := tx.Model(&ticket_token).Where("id = ?", id).Update("status","inactive").Error; err != nil {
		tx.Rollback()
		return "Error Occurred!", err, id
	}

	db.Where("id = ?", id).Delete(&ticket_token)
	if err := tx.Find(&ticket_token).Error; err != nil {
		tx.Rollback()
		return "Error Occurred!", err, id
	}

	return "Successfully Created!", err, id
}