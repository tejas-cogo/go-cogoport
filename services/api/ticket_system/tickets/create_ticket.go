package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	// timings "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_default_timings"
	// reviewers "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_reviewers"
	// "fmt"
	// "time"
)

type TicketService struct {
	Ticket models.Ticket
}

func CreateTicket(ticket models.Ticket) models.Ticket {
	db := config.GetDB()
	// result := map[string]interface{}{}

	// t1 := time.Now()

	// var filters models.TicketDefaultTiming

	// filters.TicketType =  ticket.Type
	// filters.TicketPriority =  ticket.Priority
	// filters.Status =  "active"

	// ticket_default_timing := timings.ListTicketDefaultTiming(filters)

	// for _, u := range ticket_default_timing {
		
	// 	fmt.Println(u.ID)
	// 	ticket.Tat = u.Tat
	// 	ticket.ExpiryDate = u.ExpiryDuration
	// }
	// t2 := data[tat]
	// t2 = t1+t2
	// t1 = t1.Add(t2)
	// ticket.Tat = t1

	db.Create(&ticket)

	// reviewers.CreateTicketReviewer(ticket)

	return ticket
}