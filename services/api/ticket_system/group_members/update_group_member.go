package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func UpdateGroupMember(body models.GroupMember) models.GroupMember {
	db := config.GetDB()
	var group_member models.GroupMember

	if body.ID != 0 {
		db.Where("id = ?", body.ID)
	}
	db.Find(&group_member)

	if body.ActiveTicketCount != 0 {
		group_member.ActiveTicketCount = body.ActiveTicketCount
	}
	if body.HierarchyLevel != 0 {
		group_member.HierarchyLevel = body.HierarchyLevel
	}
	if body.Status != "" {
		group_member.Status = body.Status
	}
	if body.TicketUserID != 0 {
		group_member.HierarchyLevel = body.HierarchyLevel
	}

	db.Save(&group_member)
	return group_member
}
