package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	"errors"
)

func DeleteRole(id uint) (uint,error) {
	db := config.GetDB()
	tx := db.Begin()
	var err error

	var role models.Role
	var ticket_user models.TicketUser

	if err := tx.Model(&role).Where("id = ?", id).Update("status", "inactive").Error; err != nil {
		tx.Rollback()
		return id, errors.New("Error Occurred!")
	}

	if err := tx.Model(&ticket_user).Where("role_id = ?", id).Update("role_id", 1).Error; err != nil {
		tx.Rollback()
		return id, errors.New("Error Occurred!")
	}

	if err := tx.Where("id = ?", id).Delete(&role).Error; err != nil {
		tx.Rollback()
		return id, errors.New("Error Occurred!")
	}

	tx.Commit()

	return id, err
}
