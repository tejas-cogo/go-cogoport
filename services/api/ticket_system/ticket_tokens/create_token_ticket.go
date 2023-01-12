package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	tickets "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/tickets"
	"time"
)

func CreateTokenTicket(token string, ticket models.Ticket) models.TicketToken {
	db := config.GetDB()

	var ticket_token models.TicketToken

	db.Where("ticket_token = ?",token)

	db.Find(&ticket_token)

	if ticket_token.ExpiryDate != time.Now(){
		 
		ticket.TicketUserID = ticket_token.TicketUserID
		ticket_data := tickets.CreateTicket(ticket)
		ticket_token.TicketID = ticket_data.ID
		db.Save(&ticket_token)
	}

	return ticket_token
}