package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	"gorm.io/gorm"
	activities "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_activities"
	spectators "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_spectators"
	reviewers "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_reviewers"
)

func ListTicketDetail(filters models.TicketDetail) (models.TicketDetail, *gorm.DB) {
	db := config.GetDB()

	var ticket_detail models.TicketDetail

	var ticket_reviewer models.TicketReviewer
	ticket_reviewer.TicketID = filters.TicketID 
	ticket_detail.TicketReviewer,_ = reviewers.ListTicketReviewer(ticket_reviewer)

	var ticket_spectator models.TicketSpectator
	ticket_spectator.TicketID = filters.TicketID
	ticket_detail.TicketSpectator,_ = spectators.ListTicketSpectator(ticket_spectator)

	var ticket_activity models.TicketActivity
	ticket_activity.TicketID = filters.TicketID
	ticket_detail.TicketActivity,_ = activities.ListTicketActivity(ticket_activity)

	var ticket models.Ticket
	ticket.ID = filters.TicketID
	ticket_detail.Ticket,_ = ListTicket(ticket)

	return ticket_detail, db
}
