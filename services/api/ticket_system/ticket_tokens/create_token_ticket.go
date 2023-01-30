package ticket_system

import (
	"time"

	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	tickets "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/tickets"
)

func CreateTokenTicket(ticket_token models.TicketToken) string {
	db := config.GetDB()

	db.Where("ticket_token = ?", ticket_token.TicketToken)

	db.Find(&ticket_token)

	today := time.Now()

	if today.Before(ticket_token.ExpiryDate) && ticket_token.Status != "inactive" {

		var ticket models.Ticket

		ticket.Source = ticket_token.Source
		ticket.Type = ticket_token.Type
		ticket.Category = ticket_token.Category
		ticket.Subcategory = ticket_token.Subcategory
		ticket.Description = ticket_token.Description
		ticket.Priority = ticket_token.Priority
		ticket.Tags = ticket_token.Tags
		ticket.Data = ticket_token.Data
		ticket.NotificationPreferences = ticket_token.NotificationPreferences
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
