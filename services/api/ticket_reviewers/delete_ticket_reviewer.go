package api

import (
	"errors"

	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func DeleteTicketReviewer(id uint) (uint, error) {
	db := config.GetDB()
	tx := db.Begin()
	var err error

	var ticket_reviewer models.TicketReviewer

	if err := tx.Model(&ticket_reviewer).Where("id = ?", id).Delete(&ticket_reviewer).Error; err != nil {
		tx.Rollback()
		return id, errors.New(err.Error())
	}

	if err := tx.Where("id = ?", id).Delete(&ticket_reviewer).Error; err != nil {
		tx.Rollback()
		return id, errors.New(err.Error())
	}

	tx.Commit()

	return id, err
}
