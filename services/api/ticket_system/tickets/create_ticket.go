package ticket_system

import (
	"errors"
	"fmt"
	"time"
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	audits "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_audits"
	reviewers "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_reviewers"
	helpers "github.com/tejas-cogo/go-cogoport/services/helpers"
)

type TicketService struct {
	Ticket models.Ticket
}

func CreateTicket(ticket models.Ticket) (models.Ticket, error) {
	db := config.GetDB()

	tx := db.Begin()
	var err error

	var ticket_user []models.TicketUser
	var ticket_default_type models.TicketDefaultType
	var ticket_default_timing models.TicketDefaultTiming

	if ticket.TicketUserID == 0 {
		if err := tx.Where("system_user_id = ? ", ticket.PerformedByID).Find(&ticket_user).Error; err != nil {
			tx.Rollback()
			return ticket, errors.New("System User Not Found")
		}
		if ticket_user == nil {
			return ticket, errors.New("System User Not Found")
		}
		ticket.TicketUserID = ticket_user[0].ID
	}

	if err := tx.Where("ticket_type = ? and status = ? ", ticket.Type, "active").First(&ticket_default_type).Error; err != nil {
		tx.Rollback()
		return ticket, errors.New("ticket_default_type User Not Found")
	}

	if erro := tx.Where("ticket_default_type_id = ? and status = ?", ticket_default_type.ID, "active").First(&ticket_default_timing).Error; erro != nil {
		if err := tx.Where("ticket_default_type_id = ?", 1).First(&ticket_default_timing).Error; err != nil {
			tx.Rollback()
			return ticket, errors.New("Default Timing had issue!")
		}
	}

	// fmt.Println("hjfxfghv")
	// filters.TicketDefaultTiming.TicketType = "default"
	// ticket_default_timing, err = timings.ListTicketDefaultTiming(filters.TicketDefaultTiming
	// filters.TicketDefaultTiming.TicketType = ticket.Type
	// // filters.TicketDefaultTiming.TicketPriority =
	// filters.TicketDefaultTiming.Status = "active"

	// ticket_default_timing, err := timings.ListTicketDefaultTiming(filters.TicketDefaultTiming)

	fmt.Println("rfcds", ticket_default_timing, "gfvdc")

	ticket.Priority = ticket_default_timing.TicketPriority
	ticket.Tat = ticket_default_timing.Tat
	ticket.ExpiryDate = time.Now()

	Duration := helpers.GetDuration(ticket_default_timing.ExpiryDuration)

	ticket.ExpiryDate = ticket.ExpiryDate.Add(time.Hour * time.Duration(Duration))

	ticket.Status = "unresolved"

	stmt := validate(ticket)
	if stmt != "validated" {
		return ticket, errors.New(stmt)
	}

	if err := tx.Create(&ticket).Error; err != nil {
		tx.Rollback()
		return ticket, errors.New("Ticket couldn't be created")
	}

	audits.CreateAuditTicket(ticket, db)

	var ticket_activity models.TicketActivity
	ticket_activity.TicketID = ticket.ID
	ticket_activity.TicketUserID = ticket.TicketUserID
	ticket_activity.Type = "ticket_created"
	ticket_activity.Status = "unresolved"

	if err := tx.Create(&ticket_activity).Error; err != nil {
		tx.Rollback()
		return ticket, errors.New("Activity couldn't be created")
	}
	ticket, err = reviewers.CreateTicketReviewer(ticket)
	if err != nil {
		return ticket, err
	}
	tx.Commit()

	return ticket, err

}

func validate(ticket models.Ticket) string {
	if ticket.Type == "" {
		return ("Ticket Type Is Required!")
	}
	if ticket.NotificationPreferences == nil {
		return ("Notification Preferences Is Required!")
	}
	if ticket.Tat == "" {
		return ("Tat couldn't be set!")
	}
	if ticket.ExpiryDate == time.Now() {
		return ("Expiry Date  couldn't be set!")
	}

	return ("validated")
}
