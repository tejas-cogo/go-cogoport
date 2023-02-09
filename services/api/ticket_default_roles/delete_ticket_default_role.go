package api

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	"errors"
)

func DeleteTicketDefaultRole(id uint) (uint,error){
	db := config.GetDB()
	tx := db.Begin()
	var err error

	var ticket_default_role models.TicketDefaultRole

	if err := tx.Model(&ticket_default_role).Where("id = ?", id).Update("status","inactive").Error; err != nil {
		tx.Rollback()
		return id, errors.New(err.Error())
	}

	if err := tx.Where("id = ?", id).Delete(&ticket_default_role).Error; err != nil {
		tx.Rollback()
		return id, errors.New(err.Error())
	}

	tx.Commit()

	return id, err
}