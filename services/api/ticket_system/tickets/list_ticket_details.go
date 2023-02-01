package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/models"
	activities "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_activities"
	reviewers "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_reviewers"
	spectators "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_spectators"
)

func ListTicketDetail(filters models.TicketExtraFilter) models.TicketDetail {

	var ticket_detail models.TicketDetail

	ticket_data, _ := ListTicket(filters)
	for _, u := range ticket_data {
		ticket_detail.Ticket = u
	}

	var ticket_reviewer models.TicketReviewer
	ticket_reviewer.TicketID = filters.ID
	ticket_reviewer_data, _ := reviewers.ListTicketReviewer(ticket_reviewer)
	for _, u := range ticket_reviewer_data {
		ticket_detail.TicketReviewer = u
	}

	var ticket_spectator models.TicketSpectator
	ticket_spectator.TicketID = filters.ID
	ticket_spectator_data, _, _ := spectators.ListTicketSpectator(ticket_spectator)
	for _, u := range ticket_spectator_data {
		ticket_detail.TicketSpectator = u
	}

	var ticket_activity models.TicketActivity
	ticket_activity.TicketID = filters.ID
	ticket_detail.TicketActivity, _, _ = activities.ListTicketActivity(ticket_activity)

	return ticket_detail
}
