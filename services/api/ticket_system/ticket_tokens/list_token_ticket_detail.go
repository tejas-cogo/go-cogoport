package ticket_system

import (
	"errors"

	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	reviewers "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_reviewers"
	spectators "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_spectators"
	tickets "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/tickets"
)

func ListTokenTicketDetail(token_filter models.TokenFilter) (models.TicketDetail, error) {

	var ticket_detail models.TicketDetail
	var ticket_token models.TicketToken
	var filters models.TicketExtraFilter
	var err error

	db := config.GetDB()

	tx := db.Begin()

	if err = tx.Where("ticket_token = ? and status= ?", token_filter.TicketToken, "utilized").First(&ticket_token).Error; err != nil {
		tx.Rollback()
		return ticket_detail, errors.New("Token Not Found!")
	}

	tx.Commit()

	if ticket_token.ID > 0 {
		filters.ID = ticket_token.TicketID
	}

	ticket_data, _ := tickets.ListTicket(filters)
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

	return ticket_detail, err
}
