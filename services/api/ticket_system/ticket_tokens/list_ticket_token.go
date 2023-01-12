package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func ListTicketToken() []models.TicketToken {
	db := config.GetDB()

	var ticket_token []models.TicketToken

	db.Find(&ticket_token)

	return ticket_token
}