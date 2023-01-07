package ticket_system

import (
	"github.com/ChandelShikha/go-cogoport/config"
	"github.com/ChandelShikha/go-cogoport/models"
)

type GroupMemberService struct {
	GroupMember models.GroupMember
}

func CreateGroupMember(group_member models.GroupMember) models.GroupMember {
	db := config.GetDB()
	// result := map[string]interface{}{}
	db.Create(&group_member)
	return group_member
}