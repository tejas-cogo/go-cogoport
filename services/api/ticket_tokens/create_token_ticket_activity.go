package api

import (
	"errors"

	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	activities "github.com/tejas-cogo/go-cogoport/services/api/ticket_activities"
)

func CreateTokenTicketActivity(token_activity models.TokenActivity) (models.TicketToken, error) {
	db := config.GetDB()
	tx := db.Begin()
	var err error

	var ticket_token models.TicketToken

	if err := tx.Where("ticket_token = ?", token_activity.TicketToken).First(&ticket_token).Error; err != nil {
		tx.Rollback()
		return ticket_token, errors.New(err.Error())
	}

	var body models.Filter

	var ticket []uint
	var ticket_user models.TicketUser

	ticket = append(ticket, ticket_token.TicketID)

	db.Where("id = ?", ticket_token.TicketUserID).First(&ticket_user)

	body.Activity.TicketID = ticket
	body.TicketActivity.UserID = ticket_user.SystemUserID

	body.TicketActivity.Description = token_activity.Description
	body.TicketActivity.Data = token_activity.Data
	body.TicketActivity.Status = token_activity.Status
	body.TicketActivity.Type = token_activity.Type

	_, err = activities.CreateTicketActivity(body)

	if err != nil {
		return ticket_token, err
	}

	tx.Commit()
	return ticket_token, err

}
