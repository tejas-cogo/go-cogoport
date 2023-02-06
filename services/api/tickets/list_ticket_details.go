package api

import (
	"errors"
	"fmt"

	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func ListTicketDetail(filters models.TicketExtraFilter) (models.TicketDetail, error) {

	var ticket_detail models.TicketDetail
	db := config.GetDB()
	tx := db.Begin()
	var err error
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

	fmt.Println(ticket_detail)

	tx.Commit()
	return ticket_detail, err
}
