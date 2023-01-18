package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

type GroupMemberService struct {
	GroupMember models.GroupMember
}

func CreateGroupMember(group_members models.CreateGroupMember) string {
	// result := map[string]interface{}{}
	db := config.GetDB()
	for _, u := range group_members.TicketUserID {
		var group_member models.GroupMember
		group_member.HierarchyLevel = group_members.HierarchyLevel
		group_member.GroupID = group_members.GroupID
		group_member.Status = "active"
		group_member.TicketUserID = u
		db.Create(&group_member)
	}

	return "successfully message"
}
