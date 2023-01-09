package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	group_member "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/group_members"
	ticket_default_group "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_default_groups"
)

type TicketReviewerService struct {
	TicketReviewer models.TicketReviewer
}

func CreateTicketReviewer(ticket models.Ticket) models.TicketReviewer {
	db := config.GetDB()
	// result := map[string]interface{}{}

	var ticket_default_group_filters models.TicketDefaultGroup
	var ticket_reviewer models.TicketReviewer

	ticket_reviewer.TicketId = ticket.ID
	ticket_reviewer.Status = "active"
	ticket_reviewer.TicketUserId = ticket.TicketUserId

	ticket_default_group_filters.TicketType = ticket.Type
	ticket_default_group := ticket_default_group.ListTicketDefaultGroup(ticket_default_group_filters)
	for _, u := range ticket_default_group {
		ticket_reviewer.GroupId = u.GroupId
		break
	}

	var group_member_filters models.GroupMember

	group_member_filters.GroupId = ticket_reviewer.GroupId
	group_member_data := group_member.ListGroupMember(group_member_filters)

	for _, u := range group_member_data {
		ticket_reviewer.GroupMemberId = u.ID
		group_member_filters.ActiveTicketCount += 1
		group_member_filters.TicketUserId = ticket_reviewer.TicketUserId
		group_member.UpdateGroupMember(u.ID, group_member_filters)
		break
	}

	db.Create(&ticket_reviewer)
	return ticket_reviewer
}
