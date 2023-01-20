package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/models"
	activities "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_activities"
	ticketdefaulttiming "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_default_timings"
	reviewers "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_reviewers"
	spectators "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_spectators"
)

func ListTicketDetail(filters models.TicketDetail) models.TicketDetail {

	var ticket_detail models.TicketDetail
	var ticket models.Ticket
	var sort models.Sort
	ticket.ID = filters.TicketID
	ticket_data, _ := ListTicket(ticket, sort)
	for _, u := range ticket_data {
		ticket_detail.Ticket = u
	}

	var ticket_reviewer models.TicketReviewer
	ticket_reviewer.TicketID = filters.TicketID
	ticket_reviewer_data, _ := reviewers.ListTicketReviewer(ticket_reviewer)
	for _, u := range ticket_reviewer_data {
		ticket_detail.TicketReviewer = u
	}

	var ticket_spectator models.TicketSpectator
	ticket_spectator.TicketID = filters.TicketID
	ticket_spectator_data, _ := spectators.ListTicketSpectator(ticket_spectator)
	for _, u := range ticket_spectator_data {
		ticket_detail.TicketSpectator = u
	}

	var ticket_activity models.TicketActivity
	ticket_activity.TicketID = filters.TicketID
	ticket_detail.TicketActivity, _ = activities.ListTicketActivity(ticket_activity)

	var ticket_default_timing models.TicketDefaultTiming
	var ticket_default_timings []models.TicketDefaultTiming
	ticket_default_timing.TicketType = ticket_detail.Ticket.Type
	ticket_default_timings, _ = ticketdefaulttiming.ListTicketDefaultTiming(ticket_default_timing)
	for _, u := range ticket_default_timings {
		ticket_detail.Priority = u.TicketPriority
	}

	return ticket_detail
}
