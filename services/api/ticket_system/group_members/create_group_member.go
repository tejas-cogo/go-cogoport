package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

type GroupMemberService struct {
	GroupMember models.GroupMember
}

func CreateGroupMember(group_members models.CreateGroupMember) (string, error) {
	db := config.GetDB()
	tx := db.Begin()
	var err error

	if len(group_members.TicketUserID) != 0 {
		for _, u := range group_members.TicketUserID {
			var group_member models.GroupMember
			group_member.HierarchyLevel = group_members.HierarchyLevel
			group_member.GroupID = group_members.GroupID
			group_member.Status = "active"
			group_member.TicketUserID = u
			stmt := validate(group_member)
			if stmt != "validated" {

				return stmt, err
			}
			if err := tx.Create(&group_member).Error; err != nil {
				tx.Rollback()
				return "Error Occurred!", err
			}

		}
	} else {
		return ("User Is Required!"), err
	}

	tx.Commit()

	return "Successfully Created!", err
}

func validate(group_member models.GroupMember) string {

	if group_member.HierarchyLevel == 0 {
		return ("Hierarchy Level Is Required!")
	}
	if group_member.GroupID == 0 {
		return ("Group Is Required!")
	}

	return ("validated")
}
