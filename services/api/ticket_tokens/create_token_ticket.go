package api

import (
	"errors"
	"time"

	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	tickets "github.com/tejas-cogo/go-cogoport/services/api/tickets"

	validations "github.com/tejas-cogo/go-cogoport/services/validations"
)

func CreateTokenTicket(token_filter models.TokenFilter) (models.TicketToken, error) {
	db := config.GetDB()
	tx := db.Begin()
	var err error
	var ticket_token models.TicketToken

	if err := tx.Where("ticket_token = ? AND status = ?", token_filter.TicketToken, "active").First(&ticket_token).Error; err != nil {
		db.Where("ticket_token = ?", token_filter.TicketToken).First(&ticket_token)
		var err error
		tx.Commit()
		return ticket_token, err
	}

	today := time.Now()

	if ticket_token.TicketID != 0 {
		return ticket_token, errors.New("ticket already created")
	}

	if today.Before(ticket_token.ExpiryDate) {

		var ticket models.Ticket

		ticket.Source = token_filter.Source
		ticket.Type = token_filter.Type
		ticket.Description = token_filter.Type
		ticket.TicketUserID = ticket_token.TicketUserID
		ticket.UserType = "ticket_user"

		stmt := validations.ValidateTokenTicket(ticket)
		if stmt != "validated" {
			return ticket_token, errors.New(stmt)
		}

		ticket_data, err := tickets.CreateTicket(ticket)

		if err != nil {
			return ticket_token, err
		}

		ticket_token.TicketID = ticket_data.ID
		ticket_token.Status = "misc_ticket_created"

		if err := tx.Save(&ticket_token).Error; err != nil {
			tx.Rollback()
			return ticket_token, errors.New(err.Error())
		}

		tx.Commit()

	} else {
		DeleteTicketToken(ticket_token.ID)
		tx.Commit()
	}
	return ticket_token, err
}
