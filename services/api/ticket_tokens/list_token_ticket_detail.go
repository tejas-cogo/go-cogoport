package api

import (
	"errors"

	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
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

	if ticket_token.ID > 0 {
		filters.ID = ticket_token.TicketID
	}

	var ticket models.Ticket
	var ticket_reviewer models.TicketReviewer
	// var ticket_spectator models.TicketSpectator

	if err := tx.Where("id = ?", filters.ID).First(&ticket).Error; err != nil {
		tx.Rollback()
		return ticket_detail, errors.New(err.Error())
	}
	ticket_detail.TicketID = ticket.ID
	ticket_detail.Ticket = ticket

	if err := tx.Where("ticket_id = ? and status = ?", filters.ID, "active").First(&ticket_reviewer).Error; err != nil {
		tx.Rollback()
		return ticket_detail, errors.New(err.Error())
	}
	ticket_detail.TicketReviewerID = ticket_reviewer.ID
	ticket_detail.TicketReviewer = ticket_reviewer

	// db.Where("ticket_id = ? and status = ?",filters.ID,"active").First(&ticket_spectator)
	// ticket_detail.TicketSpectatorID = ticket_spectator.ID
	// ticket_detail.TicketSpectator = ticket_spectator

	tx.Commit()

	return ticket_detail, err
}
