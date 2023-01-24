package ticket_system

import (
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

	tx := db.Begin()

	var filters models.Filter
	var ticket_reviewer models.TicketReviewer

	ticket_reviewer.TicketID = body.ID

	filters.TicketDefaultGroup.TicketType = body.Type
	filters.TicketDefaultGroup.Status = "active"
	default_group, err := defaultgroup.ListTicketDefaultGroup(filters.TicketDefaultGroup)
	if err != nil {
		return "Default Group had issue!", err
	} else if len(default_group) == 0 {
		var default_group []models.TicketDefaultGroup
		if err := tx.Where("ticket_type = ? and status = ?", "others", "active").Find(&default_group).Error; err != nil {
			tx.Rollback()
			return "Default Group couldn't be found", err
		}
	}

	for _, u := range default_group {
		ticket_reviewer.GroupID = u.GroupID
		filters.GroupMember.GroupID = u.GroupID
		filters.GroupMember.Status = "active"
		group_member, _ := groupmember.ListGroupMember(filters.GroupMember)
		for _, v := range group_member {
			ticket_reviewer.GroupMemberID = v.ID
			ticket_reviewer.TicketUserID = v.TicketUserID
			filters.GroupMember.ID = v.ID
			ticket_reviewer.Status = "active"

			if err := tx.Create(&ticket_reviewer).Error; err != nil {
				tx.Rollback()
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
	ticket_activity.Type = "Reviewer Assigned"
	ticket_activity.Status = "assigned"

	if err := tx.Create(&ticket_activity).Error; err != nil {
		tx.Rollback()
		return "Reviewer Assigned Activity couldn't be created", err
	}
	return "Successfully Reviewer Assigned", err
}
