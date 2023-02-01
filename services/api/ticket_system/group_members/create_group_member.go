package ticket_system

import (
	"errors"
	validations "github.com/tejas-cogo/go-cogoport/services/validations"
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

type GroupMemberService struct {
	GroupMember models.GroupMember
}

func CreateGroupMember(group_members models.CreateGroupMember) (models.GroupMember, error) {
	db := config.GetDB()
	tx := db.Begin()
	var err error
	var group_member models.GroupMember

	if len(group_members.TicketUserID) != 0 {
		for _, u := range group_members.TicketUserID {
			var group_member models.GroupMember
			group_member.HierarchyLevel = group_members.HierarchyLevel
			group_member.GroupID = group_members.GroupID
			group_member.Status = "active"
			group_member.TicketUserID = u
			stmt := validations.validate_group_member(group_member)
			if stmt != "validated" {

				return group_member, errors.New(stmt)
			}
			if err := tx.Create(&group_member).Error; err != nil {
				tx.Rollback()
				return group_member, errors.New("Error Occurred!")
			}

		}
	} else {
		return group_member, errors.New("User Is Required!")
	}

	tx.Commit()

	return group_member, err
}
