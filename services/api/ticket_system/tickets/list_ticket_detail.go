package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/models"
	activities "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_activities"
	reviewers "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_reviewers"
	spectators "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_spectators"
	"gorm.io/gorm"
)

func ListTicketDetail(filters models.Ticket) (models.TicketDetail, *gorm.DB) {

	var ticket models.Ticket
	var db *gorm.DB

	var ticket_reviewer models.TicketReviewer
	ticket_reviewer.TicketID = filters.ID
	ticket.TicketDetail.TicketReviewer, _ = reviewers.ListTicketReviewer(ticket_reviewer)

	var ticket_spectator models.TicketSpectator
	ticket_spectator.TicketID = filters.ID
	ticket.TicketDetail.TicketSpectator, _ = spectators.ListTicketSpectator(ticket_spectator)

	var ticket_activity models.TicketActivity
	ticket_activity.TicketID = filters.ID
	ticket.TicketDetail.TicketActivity, _ = activities.ListTicketActivity(ticket_activity)

	ticket.ID = filters.ID
	ticket.TicketDetail.Ticket, db = ListTicket(ticket)

	return ticket.TicketDetail, db
}
