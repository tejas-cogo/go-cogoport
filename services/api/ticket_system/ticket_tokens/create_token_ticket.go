package ticket_system

import (
	"time"

	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	tickets "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/tickets"
)

func CreateTokenTicket(token string, ticket models.Ticket) string {
	db := config.GetDB()

	var ticket_token models.TicketToken

	db.Where("ticket_token = ?", token)

	db.Find(&ticket_token)

	today := time.Now()

	if today.Before(ticket_token.ExpiryDate) && ticket_token.Status != "inactive" {

		ticket.TicketUserID = ticket_token.TicketUserID
		ticket_data, mesg, _ := tickets.CreateTicket(ticket)

		if mesg == "Successfully Created!" {
			ticket_token.TicketID = ticket_data.ID
		} else {
			return mesg
		}

		ticket_token.Status = "used"
		db.Save(&ticket_token)
	} else {
		DeleteTicketToken(ticket_token.ID)
	}
	return "Successfully Created!"
}
