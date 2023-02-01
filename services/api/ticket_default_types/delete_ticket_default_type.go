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
	var ticket_default_group models.TicketDefaultGroup
	var ticket_default_timing models.TicketDefaultTiming

	if err := tx.Model(&ticket_default_type).Where("id = ?", id).Update("status", "inactive").Error; err != nil {
		tx.Rollback()
		return id, errors.New("Cannot update ticket default type status!")
	}

	if err := tx.Model(&ticket_default_group).Where("ticket_default_type_id = ?", id).Update("status", "inactive").Error; err != nil {
		tx.Rollback()
		return id, errors.New("Cannot update ticket default group status!")
	}

	if err := tx.Where("ticket_default_type_id = ?", id).Delete(&ticket_default_group).Error; err != nil {
		tx.Rollback()
		return id, errors.New("Cannot delete ticket default group!")
	}

	if err := tx.Model(&ticket_default_timing).Where("ticket_default_type_id = ?", id).Update("status", "inactive").Error; err != nil {
		tx.Rollback()
		return id, errors.New("Cannot update ticket default time status!")
	}

	if err := tx.Where("ticket_default_type_id = ?", id).Delete(&ticket_default_timing).Error; err != nil {
		tx.Rollback()
		return id, errors.New("Cannot delete ticket default timing!")
	}

	if err := tx.Where("id = ?", id).Delete(&ticket_default_type).Error; err != nil {
		tx.Rollback()
		return id, errors.New("Cannot deleted ticket default type!")
	}

	tx.Commit()

	return id, err
}
