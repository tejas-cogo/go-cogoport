package api

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	"errors"
)

func DeleteTicketDefaultType(id uint) (uint,error) {
	db := config.GetDB()
	tx := db.Begin()
	var err error

	var ticket_default_type models.TicketDefaultType
	var ticket_default_role models.TicketDefaultRole
	var ticket_default_timing models.TicketDefaultTiming

	if err := tx.Model(&ticket_default_type).Where("id = ?", id).Update("status", "inactive").Error; err != nil {
		tx.Rollback()
		return id, errors.New(err.Error())
	}

	if err := tx.Model(&ticket_default_role).Where("ticket_default_type_id = ?", id).Update("status", "inactive").Error; err != nil {
		tx.Rollback()
		return id, errors.New(err.Error())
	}

	if err := tx.Where("ticket_default_type_id = ?", id).Delete(&ticket_default_role).Error; err != nil {
		tx.Rollback()
		return id, errors.New(err.Error())
	}

	if err := tx.Model(&ticket_default_timing).Where("ticket_default_type_id = ?", id).Update("status", "inactive").Error; err != nil {
		tx.Rollback()
		return id, errors.New(err.Error())
	}

	if err := tx.Where("ticket_default_type_id = ?", id).Delete(&ticket_default_timing).Error; err != nil {
		tx.Rollback()
		return id, errors.New(err.Error())
	}

	if err := tx.Where("id = ?", id).Delete(&ticket_default_type).Error; err != nil {
		tx.Rollback()
		return id, errors.New(err.Error())
	}

	tx.Commit()

	return id, err
}
