package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func DeleteGroup(id uint) uint {
	db := config.GetDB()

	var group models.Group
	var group_member models.GroupMember

	// db.Where("id = ?", id).First(&group)
	// group.Status = "inactive"
	// db.Save(&group)
	db.Model(&group).Where("id = ?", id).Update("status", "inactive")

	db.Model(&group_member).Where("group_id = ? ", id).Update("status", "inactive")
	db.Where("id = ?", id).Delete(&group_member)

	db.Where("id = ?", id).Delete(&group)

	return id
}
