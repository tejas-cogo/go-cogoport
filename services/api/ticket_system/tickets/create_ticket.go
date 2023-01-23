package ticket_system

import (
	"strconv"
	"strings"
	"time"

	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	audits "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_audits"
	timings "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_default_timings"
	reviewers "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_reviewers"
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

func CreateTicket(ticket models.Ticket) (models.Ticket, error) {
	db := config.GetDB()
	// result := map[string]interface{}{}

	tx := db.Begin()
	var err error

	var filters models.Filter

	var ticket_user models.TicketUser

	if ticket.TicketUserID == 0 {
		if err := tx.Where("system_user_id = ? ", ticket.PerformedByID).Find(&ticket_user).Error; err != nil {
			tx.Rollback()
			return ticket, err
		}
		ticket.TicketUserID = ticket_user.ID
	}

	filters.TicketDefaultTiming.TicketType = ticket.Type
	// filters.TicketDefaultTiming.TicketPriority = ticket.Priority
	filters.TicketDefaultTiming.Status = "active"

	ticket_default_timing, _ := timings.ListTicketDefaultTiming(filters.TicketDefaultTiming)

	for _, u := range ticket_default_timing {

		ticket.Tat = u.Tat
		ticket.ExpiryDate = time.Now()

		Duration := GetDuration(u.ExpiryDuration)

		ticket.ExpiryDate = ticket.ExpiryDate.Add(time.Hour * time.Duration(Duration))
		break
	}
	ticket.Status = "unresolved"

	if err := tx.Create(&ticket).Error; err != nil {
		tx.Rollback()
		return ticket, err
	}

	audits.CreateAuditTicket(ticket, db)

	var ticket_activity models.TicketActivity
	ticket_activity.TicketID = ticket.ID
	ticket_activity.TicketUserID = ticket.TicketUserID
	ticket_activity.Type = "Ticket Created"
	ticket_activity.Status = "unresolved"

	if err := tx.Create(&ticket_activity).Error; err != nil {
		tx.Rollback()
		return ticket, err
	}

	if err := reviewers.CreateTicketReviewer(ticket); err != nil {
		tx.Rollback()
		return ticket, err
	}

	tx.Commit()

	return ticket, err

}
