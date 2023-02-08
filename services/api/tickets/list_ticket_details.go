package api

import (
	"errors"

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
	var ticket_reviewer_data models.TicketReviewerData
	// var ticket_spectator models.TicketSpectator

	if err := tx.Where("id = ?", filters.ID).First(&ticket).Error; err != nil {
		tx.Rollback()
		return ticket_detail, errors.New(err.Error())
	}
	ticket_detail.TicketID = ticket.ID
	ticket_detail.Ticket = ticket

	if err := tx.Model(&ticket_reviewer).Where("ticket_id = ? and status != ?", filters.ID, "inactive").Scan(&ticket_reviewer_data).Error; err != nil {
		tx.Rollback()
		return ticket_detail, errors.New(err.Error())
	}
	ticket_detail.TicketReviewerID = ticket_reviewer.ID
	ticket_detail.TicketReviewer = ticket_reviewer_data

	var t_user models.TicketUser
	db.Where("system_user_id = ?", ticket.UserID).First(&t_user)
	ticket_detail.TicketUser = t_user

	tx.Commit()
	return ticket_detail, err
}
