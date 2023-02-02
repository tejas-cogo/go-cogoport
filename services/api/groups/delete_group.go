package api

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	"errors"
)

func DeleteGroup(id uint) (uint,error) {
	db := config.GetDB()
	tx := db.Begin()
	var err error

	var group models.Group
	var group_member models.GroupMember

	if err := tx.Model(&group).Where("id = ?", id).Update("status", "inactive").Error; err != nil {
		tx.Rollback()
		return id, errors.New("Cannot update group details!")
	}

	if err := tx.Model(&group_member).Where("group_id = ? ", id).Update("status", "inactive").Error; err != nil {
		tx.Rollback()
		return id, errors.New("Cannot update group member details!")
	}

	if err := tx.Where("id = ?", id).Delete(&group_member).Error; err != nil {
		tx.Rollback()
		return id, errors.New("Cannot delete group member details!")
	}

	if err := tx.Where("id = ?", id).Delete(&group).Error; err != nil {
		tx.Rollback()
		return id, errors.New("Cannot update group details!")
	}

	tx.Commit()

	return id, err
}
