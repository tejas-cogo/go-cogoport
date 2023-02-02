package api

import (
	"errors"

	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func DeleteTicket(body models.Ticket) (models.Ticket, error) {
	db := config.GetDB()
	tx := db.Begin()
	var err error

	var ticket models.Ticket

	if err := tx.Model(&ticket).Where("id = ?", body.ID).Update("status", "inactive").Error; err != nil {
		tx.Rollback()
		return body, errors.New("Cannot find ticket with this id!")
	}

	if err := tx.Where("id = ?", body.ID).Delete(&ticket).Error; err != nil {
		tx.Rollback()
		return body, errors.New("Cannot delete ticket!")
	}

	tx.Commit()
	return body, err
}
