package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	"errors"
)

func DeleteTicketToken(id uint) (uint,error){
	db := config.GetDB()
	tx := db.Begin()
	var err error

	var ticket_token models.TicketToken

	if err := tx.Model(&ticket_token).Where("id = ?", id).Update("status","inactive").Error; err != nil {
		tx.Rollback()
		return id, errors.New("Error Occurred!")
	}

	db.Where("id = ?", id).Delete(&ticket_token)
	if err := tx.Find(&ticket_token).Error; err != nil {
		tx.Rollback()
		return id, errors.New("Error Occurred!")
	}

	return id, err
}