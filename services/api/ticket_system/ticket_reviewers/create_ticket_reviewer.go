package ticket_system

import (
	"fmt"

	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	groupmember "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/group_members"
	defaultgroup "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_default_groups"
)

type TicketReviewerService struct {
	TicketReviewer   models.TicketReviewer
	ReviewerActivity models.ReviewerActivity
}

func CreateTicketReviewer(body models.Ticket) (string, error) {
	db := config.GetDB()
	// result := map[string]interface{}{}

	txt := db.Begin()

	var filters models.Filter
	var ticket_reviewer models.TicketReviewer

	filters.TicketDefaultGroup.TicketType = body.Type
	filters.TicketDefaultGroup.Status = "active"
	default_group, err := defaultgroup.ListTicketDefaultGroup(filters.TicketDefaultGroup)
	if err != nil {
		return "Default Group had issue!", err
	} else if len(default_group) == 0 {
		var default_group []models.TicketDefaultGroup
		if err := txt.Where("ticket_type = ? ", "default").Find(&default_group).Error; err != nil {
			txt.Rollback()
			return "Default Group couldn't be found", err
		}
	}

	for _, u := range default_group {
		ticket_reviewer.GroupID = u.GroupID
		var group_member []models.GroupMember
		if u.GroupMemberID < 1 {
			filters.FilterGroupMember.GroupID = u.GroupID
			filters.FilterGroupMember.Status = "active"
			filters.FilterGroupMember.NotPresentTicketUserID = body.TicketUserID
			fmt.Println("group", filters.FilterGroupMember)
			group_member, _ = groupmember.ListGroupMember(filters.FilterGroupMember)
		} else {
			txt.Where("id = ?", u.GroupMemberID).Find(&group_member)
		}

		for _, v := range group_member {
			ticket_reviewer.GroupMemberID = v.ID
			ticket_reviewer.TicketUserID = v.TicketUserID
			ticket_reviewer.TicketID = body.ID
			filters.GroupMember.ID = v.ID
			ticket_reviewer.Status = "active"

			stmt := validate(ticket_reviewer)
			if stmt != "validated" {
				return stmt, err
			}
			if err := txt.Create(&ticket_reviewer).Error; err != nil {
				txt.Rollback()
				return "TicketReviewer couldn't be created", err
			}

			filters.GroupMember.ActiveTicketCount = v.ActiveTicketCount + 1
			groupmember.UpdateGroupMember(filters.GroupMember)
			break
		}
		break
	}

	var ticket_activity models.TicketActivity
	ticket_activity.TicketID = ticket_reviewer.TicketID
	ticket_activity.TicketUserID = ticket_reviewer.TicketUserID
	ticket_activity.UserType = "system"
	ticket_activity.Type = "reviewer_assigned"
	ticket_activity.Status = "assigned"

	if err := txt.Create(&ticket_activity).Error; err != nil {
		txt.Rollback()
		return "Reviewer Assigned Activity couldn't be created", err
	}

	txt.Commit()
	return "Successfully Reviewer Assigned", err
}

func validate(ticket_reviewer models.TicketReviewer) string {
	if ticket_reviewer.GroupMemberID == 0 {
		return ("Group Member Is Required!")
	}

	if ticket_reviewer.GroupID == 0 {
		return ("Group Is Required!")
	}

	if ticket_reviewer.TicketID == 0 {
		return ("Ticket Is Required!")
	}

	if ticket_reviewer.TicketUserID == 0 {
		return ("Ticket User Is Required!")
	}

	return ("validated")
}
