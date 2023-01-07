package ticket_system

import (
	"github.com/ChandelShikha/go-cogoport/config"
	"github.com/ChandelShikha/go-cogoport/models"
)

func DeleteGroupMember(id uint) uint{
	db := config.GetDB()

	var group_member models.GroupMember

	db.Model(&group_member).Where("id = ?", id).Update("status","inactive")

	db.Where("id = ?", id).Delete(&group_member)

	return id
}