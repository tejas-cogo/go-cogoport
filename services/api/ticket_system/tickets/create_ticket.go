package ticket_system

import (
	"fmt"
	"time"

	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	activities "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_activities"
	timings "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_default_timings"
	reviewers "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_reviewers"
)

type TicketService struct {
	Ticket models.Ticket
}

func CreateTicket(ticket models.Ticket) models.Ticket {
	db := config.GetDB()
	// result := map[string]interface{}{}

	t1 := time.Now()

	var filters models.TicketDefaultTiming

	filters.TicketType = ticket.Type
	filters.TicketPriority = ticket.Priority
	filters.Status = "active"

	ticket_default_timing := timings.ListTicketDefaultTiming(filters)

	for _, u := range ticket_default_timing {
		fmt.Println(u.ID)
		ticket.Tat = u.Tat
		ticket.ExpiryDate = t1.Add(u.ExpiryDuration)
	}

	db.Create(&ticket)

	reviewers.CreateTicketReviewer(ticket)

	var ticket_activity models.TicketActivity
	ticket_activity.TicketID = ticket.ID
	ticket_activity.TicketUserID = ticket.TicketUserID
	activities.CreateTicketActivity(ticket_activity)

	return ticket

}