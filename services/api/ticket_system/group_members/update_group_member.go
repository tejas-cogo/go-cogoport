package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func UpdateGroupMember(filters models.Filter) models.GroupMember {
	db := config.GetDB()
	var group_member models.GroupMember

	group_member_filters := filters.GroupMember

	if group_member_filters.TicketUserID != 0 {
		db.Where("ticket_user_id = ?", group_member_filters.TicketUserID)
	}
	if group_member_filters.ID != 0 {
		db.Where("id = ?", group_member_filters.ID)
	}
	if group_member_filters.Status != "" {
		db.Where("status = ?", group_member_filters.Status)
	}
	db.Find(&group_member)

	if group_member.GroupID == 0 {
		var ticket_default_group models.TicketDefaultGroup
		ticket_default_group_filters := filters.TicketDefaultGroup
		db.Where("ticket_type = ?", ticket_default_group_filters.TicketType).First(&ticket_default_group)
		group_member.GroupID = ticket_default_group.GroupID
	}

	if group_member_filters.ActiveTicketCount != 0 {
		group_member.ActiveTicketCount = group_member_filters.ActiveTicketCount
	}
	if group_member_filters.HierarchyLevel != 0 {
		group_member.HierarchyLevel = group_member_filters.HierarchyLevel
	}
	if group_member_filters.Status != "" {
		group_member.Status = group_member_filters.Status
	}
	if group_member_filters.TicketUserID != 0 {
		group_member.HierarchyLevel = group_member_filters.HierarchyLevel
	}

	db.Save(&group_member)
	return group_member
}
