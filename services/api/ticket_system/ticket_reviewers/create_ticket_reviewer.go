package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	groupmember "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/group_members"
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
	ticket_reviewer.Status = "active"

	filters.TicketDefaultGroup.TicketType = ticket.Type
	default_group := defaultgroup.ListTicketDefaultGroup(filters.TicketDefaultGroup)
	for _, u := range default_group {
		ticket_reviewer.GroupID = u.GroupID

		filters.GroupMember.GroupID = u.GroupID
		group_member := groupmember.ListGroupMember(filters.GroupMember)
		for _, u := range group_member {
			ticket_reviewer.GroupID = u.GroupID
			ticket_reviewer.GroupMemberID = u.ID
			filters.GroupMember.ID = u.ID
			filters.GroupMember.ActiveTicketCount += 1
			groupmember.UpdateGroupMember(filters)
			break
		}
		break
	}

	db.Create(&ticket_reviewer)
	return ticket_reviewer
}
