package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func DeleteGroupMember(id uint) (string,uint) {
	db := config.GetDB()

	var group_member models.GroupMember
	var member models.GroupMember

	db.Where("id != ? and status = ?", id, "active").Find(&group_member)

	db.Where("group_id = ? and id != ? and status = ?", group_member.GroupID, id, "active").Find(&member)

	if member.ID != 0 {
		return "Cannot be deleted",id
	}
	db.Model(&group_member).Where("id = ?", id).Update("status", "inactive")

	db.Where("id = ?", id).Delete(&group_member)

	return "Successfully Deleted",id
}
