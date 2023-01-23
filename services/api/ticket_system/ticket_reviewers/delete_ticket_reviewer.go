package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func DeleteTicketReviewer(id uint) uint {
	db := config.GetDB()

	var ticket_reviewer models.TicketReviewer
	var group_member models.GroupMember

	db.Model(&ticket_reviewer).Where("id = ?", id).Update("status", "inactive")

	db.Where("id = ?", id).Delete(&ticket_reviewer)

	db.Where("id = ?", ticket_reviewer.GroupMemberID).First(&group_member)
	group_member.ActiveTicketCount -= 1
	db.Save(&group_member)

	return id
}
