package api

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	"errors"
)

func DeleteTicketDefaultTiming(id uint) (uint,error){
	db := config.GetDB()
	tx := db.Begin()
	var err error

	var ticket_default_timing models.TicketDefaultTiming

	if err := tx.Model(&ticket_default_timing).Where("id = ?", id).Update("status","inactive").Error; err != nil {
		tx.Rollback()
		return id, errors.New("Cannot update ticket default time status!")
	}

	if err := tx.Where("id = ?", id).Delete(&ticket_default_timing).Error; err != nil {
		tx.Rollback()
		return id, errors.New("Cannot delete ticket default time!")
	}

	tx.Commit()

	return id, err
}