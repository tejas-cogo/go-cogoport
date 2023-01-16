package ticket_system

import (
	"fmt"

	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	groupmember "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/group_members"
	activities "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_activities"
	defaultgroup "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_default_groups"
)

type TicketReviewerService struct {
	TicketReviewer models.TicketReviewer
}

func CreateTicketReviewer(ticket models.Ticket) models.TicketReviewer {
	db := config.GetDB()
	// result := map[string]interface{}{}

	var filters models.Filter
	var ticket_reviewer models.TicketReviewer

	ticket_reviewer.TicketID = ticket.ID

	filters.TicketDefaultGroup.TicketType = ticket.Type
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
			db.Create(&ticket_reviewer)
			filters.GroupMember.ActiveTicketCount = v.ActiveTicketCount + 1
			groupmember.UpdateGroupMember(filters)
			fmt.Print("old", v.ActiveTicketCount, "new", filters.GroupMember.ActiveTicketCount)
			break
		}
		break
	}

	filters.TicketActivity.TicketID = ticket.ID
	filters.TicketActivity.UserType = "system"
	filters.TicketActivity.Type = "Reviewer Assigned"

	activities.CreateTicketActivity(filters)

	return ticket_reviewer
}
