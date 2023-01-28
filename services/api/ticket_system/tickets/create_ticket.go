package ticket_system

import (
	"fmt"
	"time"

	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	audits "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_audits"
	timings "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_default_timings"
	reviewers "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_reviewers"
	helpers "github.com/tejas-cogo/go-cogoport/services/helpers"
)

type TicketService struct {
	Ticket models.Ticket
}

func CreateTicket(ticket models.Ticket) (models.Ticket, string, error) {
	db := config.GetDB()
	// result := map[string]interface{}{}

	tx := db.Begin()
	var err error

	var filters models.Filter
	var ticket_user []models.TicketUser

	if ticket.TicketUserID == 0 {
		if err := tx.Where("system_user_id = ? ", ticket.PerformedByID).Find(&ticket_user).Error; err != nil {
			tx.Rollback()
			return ticket, "System User Not Found", err
		}
		if ticket_user == nil {
			return ticket, "System User Not Found", err
		}
		ticket.TicketUserID = ticket_user[0].ID
	}

	filters.TicketDefaultTiming.TicketType = ticket.Type
	// filters.TicketDefaultTiming.TicketPriority = ticket.Priority
	filters.TicketDefaultTiming.Status = "active"

	ticket_default_timing, err := timings.ListTicketDefaultTiming(filters.TicketDefaultTiming)
	if err != nil {
		return ticket, "Default Timing had issue!", err
	} else if len(ticket_default_timing) == 0 {
		fmt.Println("hjfxfghv")
		filters.TicketDefaultTiming.TicketType = "default"
		ticket_default_timing, err = timings.ListTicketDefaultTiming(filters.TicketDefaultTiming)
		if err != nil || len(ticket_default_timing) == 0 {
			return ticket, "Default Timing had issue!", err
		}
	}

	fmt.Println("rfcds", ticket_default_timing, "gfvdc")

	for _, u := range ticket_default_timing {

		ticket.Tat = u.Tat
		ticket.ExpiryDate = time.Now()

		Duration := helpers.GetDuration(u.ExpiryDuration)

		ticket.ExpiryDate = ticket.ExpiryDate.Add(time.Hour * time.Duration(Duration))
	}
	ticket.Status = "unresolved"

	stmt := validate(ticket)
	if stmt != "validated" {
		return ticket, stmt, err
	}

	if err := tx.Create(&ticket).Error; err != nil {
		tx.Rollback()
		return ticket, "Ticket couldn't be created", err
	}

	audits.CreateAuditTicket(ticket, db)

	var ticket_activity models.TicketActivity
	ticket_activity.TicketID = ticket.ID
	ticket_activity.TicketUserID = ticket.TicketUserID
	ticket_activity.Type = "ticket_created"
	ticket_activity.Status = "unresolved"

	if err := tx.Create(&ticket_activity).Error; err != nil {
		tx.Rollback()
		return ticket, "Activity couldn't be created", err
	}

	if stmt, err := reviewers.CreateTicketReviewer(ticket); err != nil {
		return ticket, stmt, err
	}

	tx.Commit()

	return ticket, "Successfully Created!", err

}

func validate(ticket models.Ticket) string {
	if ticket.Type == "" {
		return ("Ticket Type Is Required!")
	}
	if ticket.NotificationPreferences == nil {
		return ("Notification Preferences Is Required!")
	}
	// if ticket.Priority == nil {
	// 	return ("Priority Is Required!")
	// }
	if ticket.Tat == "" {
		return ("Tat couldn't be set!")
	}
	if ticket.ExpiryDate == time.Now() {
		return ("Expiry Date  couldn't be set!")
	}

	return ("validated")
}
