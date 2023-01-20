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

func CreateTicketReviewer(body models.Ticket) models.TicketReviewer {
	db := config.GetDB()
	// result := map[string]interface{}{}

	var filters models.Filter
	var ticket_reviewer models.TicketReviewer

	ticket_reviewer.TicketID = body.ID

	filters.TicketDefaultGroup.TicketType = body.Type
	filters.TicketDefaultGroup.Status = "active"
	default_group, _ := defaultgroup.ListTicketDefaultGroup(filters.TicketDefaultGroup)
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
			db.Create(&ticket_reviewer)
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

	db.Create(&ticket_activity)

	return ticket_reviewer
}
