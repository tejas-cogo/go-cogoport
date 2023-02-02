package api

import (
	"github.com/tejas-cogo/go-cogoport/models"
	reviewers "github.com/tejas-cogo/go-cogoport/services/api/ticket_reviewers"
	spectators "github.com/tejas-cogo/go-cogoport/services/api/ticket_spectators"
)

func ListTicketDetail(filters models.TicketExtraFilter) models.TicketDetail {

	var ticket_detail models.TicketDetail

	ticket_data, _ := ListTicket(filters)
	for _, u := range ticket_data {
		ticket_detail.Ticket = u

		// 	// Duration := helpers.GetDuration(u.Tat)
		// 	// u.Tat = time.Now()
		// 	// ticket_detail.Ticket.Tat = Tat.Add(time.Hour * time.Duration(Duration))
	}

	var ticket_reviewer models.TicketReviewer
	ticket_reviewer.TicketID = filters.ID
	ticket_reviewer_data, _ := reviewers.ListTicketReviewer(ticket_reviewer)
	for _, u := range ticket_reviewer_data {
		ticket_detail.TicketReviewer = u
		ticket_detail.TicketReviewerID = u.ID
	}

	var ticket_spectator models.TicketSpectator
	ticket_spectator.TicketID = filters.ID
	ticket_spectator_data, _, _ := spectators.ListTicketSpectator(ticket_spectator)
	for _, u := range ticket_spectator_data {
		ticket_detail.TicketSpectator = u
		ticket_detail.TicketSpectatorID = u.ID
	}

	return ticket_detail
}
