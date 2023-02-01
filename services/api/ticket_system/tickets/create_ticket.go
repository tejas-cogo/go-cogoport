package ticket_system

import (
	"errors"
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
		if err := tx.Where("id = ?", 1).First(&ticket_default_type).Error; err != nil {
			tx.Rollback()
			return ticket, errors.New("Default Type had issue!")
		}
	}

	if erro := tx.Where("ticket_default_type_id = ? and status = ?", ticket_default_type.ID, "active").First(&ticket_default_timing).Error; erro != nil {
		if err := tx.Where("ticket_default_type_id = ?", 1).First(&ticket_default_timing).Error; err != nil {
			tx.Rollback()
			return ticket, errors.New("Default Timing had issue!")
		}
	}

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
	if ticket.Tat == "" {
		return ("Tat couldn't be set!")
	}
	if ticket.ExpiryDate == time.Now() {
		return ("Expiry Date  couldn't be set!")
	}

	return ("validated")
}
