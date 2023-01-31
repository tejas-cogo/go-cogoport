package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	"errors"
)

func DeleteGroupMember(id uint) (uint,error){
	db := config.GetDB()
	tx := db.Begin()
	var err error

	var group_member models.GroupMember


	if err := tx.Model(&group_member).Where("id = ?", id).Update("status","inactive").Error; err != nil {
		tx.Rollback()
		return id, errors.New("Error Occurred!")
	}


	if err := tx.Where("id = ?", id).Delete(&group_member).Error; err != nil {
		tx.Rollback()
		return id, errors.New("Error Occurred!")
	}

	tx.Commit()

	return id, err
}