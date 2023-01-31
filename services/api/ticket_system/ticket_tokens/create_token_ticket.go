package ticket_system

import (
	"time"

	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	tickets "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/tickets"
)

func CreateTokenTicket(token_filter models.TokenFilter) string {
	db := config.GetDB()
	var ticket_token models.TicketToken

	db.Where("ticket_token = ? AND status != ?", token_filter.TicketToken, "used").Find(&ticket_token)

	today := time.Now()

	if today.Before(ticket_token.ExpiryDate) && ticket_token.Status != "inactive" {

		var ticket models.Ticket

		ticket.Source = token_filter.Source
		ticket.Type = token_filter.Type
		ticket.Category = token_filter.Category
		ticket.Subcategory = token_filter.Subcategory
		ticket.Description = token_filter.Description
		ticket.IsUrgent = token_filter.IsUrgent
		ticket.Data = token_filter.Data
		ticket.NotificationPreferences = token_filter.NotificationPreferences
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
