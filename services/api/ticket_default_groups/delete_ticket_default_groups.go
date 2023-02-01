package api

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	"errors"
)

func DeleteTicketDefaultGroup(id uint) (uint,error){
	db := config.GetDB()
	tx := db.Begin()
	var err error

	var ticket_default_group models.TicketDefaultGroup

	if err := tx.Model(&ticket_default_group).Where("id = ?", id).Update("status","inactive").Error; err != nil {
		tx.Rollback()
		return id, errors.New("Error Occurred!")
	}

	if err := tx.Where("id = ?", id).Delete(&ticket_default_group).Error; err != nil {
		tx.Rollback()
		return id, errors.New("Error Occurred!")
	}

	tx.Commit()

	return id, err
}