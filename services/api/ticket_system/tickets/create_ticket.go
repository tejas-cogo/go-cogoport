package ticket_system

import (
	// "encoding/json"

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

	var filters models.TicketDefaultTiming

	filters.TicketType = ticket.Type
	filters.TicketPriority = ticket.Priority
	filters.Status = "active"

	ticket_default_timing := timings.ListTicketDefaultTiming(filters)

	fmt.Println(ticket_default_timing)

	for _, u := range ticket_default_timing {

		ticket.Tat = time.Hour * time.Duration(u.ExpiryDuration)
		ticket.ExpiryDate = time.Now()
		ticket.ExpiryDate = ticket.ExpiryDate.Add(time.Hour * time.Duration(u.ExpiryDuration))
		fmt.Println("start", ticket.ExpiryDate, "start")
		break
	}

	db.Create(&ticket)

	var ticket_audit models.TicketAudit

	db.Create(&ticket_audit)

	reviewers.CreateTicketReviewer(ticket)

	var ticket_activity models.TicketActivity
	ticket_activity.TicketID = ticket.ID
	ticket_activity.TicketUserID = ticket.TicketUserID
	activities.CreateTicketActivity(ticket_activity)

	return ticket

}
