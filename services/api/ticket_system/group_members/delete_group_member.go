package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func DeleteGroupMember(id uint) (string,error,uint){
	db := config.GetDB()
	tx := db.Begin()
	var err error

	var group_member models.GroupMember


	if err := tx.Model(&group_member).Where("id = ?", id).Update("status","inactive").Error; err != nil {
		tx.Rollback()
		return "Error Occurred!", err,id
	}


	if err := tx.Where("id = ?", id).Delete(&group_member).Error; err != nil {
		tx.Rollback()
		return "Error Occurred!", err,id
	}

	tx.Commit()

	return "Successfully Deleted!",err,id
}