package ticket_system

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	activities "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_activities"
	timings "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_default_timings"
	reviewers "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_reviewers"

	"gorm.io/gorm"
)

type TicketService struct {
	Ticket models.Ticket
}

func GetDuration(ExpiryDuration string) int {
	duration := strings.Split(ExpiryDuration, ":")

	durationd := strings.Split(duration[0], "d")
	durationh := strings.Split(duration[1], "h")
	durationm := strings.Split(duration[2], "m")

	d, _ := strconv.Atoi(durationd[0])
	h, _ := strconv.Atoi(durationh[0])
	m, _ := strconv.Atoi(durationm[0])

	h += m / 60
	h += d * 24

	return h

}

func CreateAuditTicket(ticket models.Ticket, db *gorm.DB) int {
	var ticket_audit models.TicketAudit

	ticket_audit.ObjectId = ticket.ID
	ticket_audit.Action = "create"
	ticket_audit.Object = "ticket"
	data, _ := json.Marshal(ticket)
	ticket_audit.Data = string(data)

	db.Create(&ticket_audit)

	return 0
}

func CreateTicket(ticket models.Ticket) models.Ticket {
	db := config.GetDB()
	// result := map[string]interface{}{}

	var filters models.Filter

	filters.TicketDefaultTiming.TicketType = ticket.Type
	filters.TicketDefaultTiming.TicketPriority = ticket.Priority
	filters.TicketDefaultTiming.Status = "active"

	ticket_default_timing, _ := timings.ListTicketDefaultTiming(filters.TicketDefaultTiming)

	for _, u := range ticket_default_timing {

		ticket.Tat = u.Tat
		ticket.ExpiryDate = time.Now()

		Duration := GetDuration(u.ExpiryDuration)

		ticket.ExpiryDate = ticket.ExpiryDate.Add(time.Hour * time.Duration(Duration))
		break
	}

	db.Create(&ticket)

	CreateAuditTicket(ticket, db)

	filters.TicketActivity.TicketID = ticket.ID
	filters.TicketActivity.TicketUserID = ticket.TicketUserID
	filters.TicketActivity.Type = "Ticket Created"

	activities.CreateTicketActivity(filters)

	reviewers.CreateTicketReviewer(ticket)

	return ticket

}
