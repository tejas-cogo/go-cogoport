package api

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	"errors"
)

func UpdateTicketActivity(body models.TicketActivity) (models.TicketActivity,error) {
	db := config.GetDB()
	tx := db.Begin()
	var err error

	var ticket_activity models.TicketActivity
	
	if err := tx.Where("id = ?", body.ID).Find(&ticket_activity).Error; err != nil {
		tx.Rollback()
		return ticket_activity, errors.New(err.Error())
	}

	if err := tx.Save(&ticket_activity).Error; err != nil {
		tx.Rollback()
		return ticket_activity, errors.New(err.Error())
	}

	tx.Commit()
	
	return ticket_activity, err
}
