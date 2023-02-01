package ticket_system

import (
	"errors"

	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func DeleteGroupMember(id uint) (uint, error) {

	db := config.GetDB()
	tx := db.Begin()
	var err error

	var group_member models.GroupMember
	var member []models.GroupMember

	if err := tx.Where("id = ? and status = ?", id, "active").Find(&group_member).Error; err != nil {
		tx.Rollback()
		return id, errors.New("Not found existed member!")

	}

	if err := tx.Where("group_id = ? and id != ? and status = ?", group_member.GroupID, id, "active").Find(&member).Error; err != nil {
		tx.Rollback()
		return id, errors.New("Error Occurred!")

	}

	if len(member) == 0 {
		return id, errors.New("Last remaining member cannot be deleted.")

	}

	if err := tx.Model(&group_member).Where("id = ?", id).Update("status", "inactive").Error; err != nil {
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
