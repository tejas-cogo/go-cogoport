package api

import (
	"errors"
	"time"

	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	audits "github.com/tejas-cogo/go-cogoport/services/api/ticket_audits"
	reviewers "github.com/tejas-cogo/go-cogoport/services/api/ticket_reviewers"
	ticket_users "github.com/tejas-cogo/go-cogoport/services/api/ticket_users"
	helpers "github.com/tejas-cogo/go-cogoport/services/helpers"
	validations "github.com/tejas-cogo/go-cogoport/services/validations"
)

type TicketService struct {
	Ticket models.Ticket
}

func CreateTicket(ticket models.Ticket) (models.Ticket, error) {
	db := config.GetDB()

	tx := db.Begin()
	var err error

	var ticket_user models.TicketUser
	var user models.User
	var ticket_default_type models.TicketDefaultType
	var ticket_default_timing models.TicketDefaultTiming

	if ticket.TicketUserID != 0 {
		if err := tx.Where("id = ? ", ticket.TicketUserID).First(&ticket_user).Error; err != nil {
			tx.Rollback()
			return ticket, errors.New(err.Error())
		}
		if ticket_user.ID == 0 {
			return ticket, errors.New("system user not found")
		}
		ticket.UserID = ticket_user.SystemUserID
	} else {
		db2 := config.GetCDB().Debug()
		var new_ticket_user models.TicketUser
		db2.Where("id = ?", ticket.UserID).First(&user)
		new_ticket_user.Email = user.Email
		new_ticket_user.SystemUserID = user.ID
		new_ticket_user.Name = user.Name
		new_ticket_user.MobileNumber = user.MobileNumber
		ticket_user, _ := ticket_users.CreateTicketUser(new_ticket_user)
		ticket.TicketUserID = ticket_user.ID
	}

	if err := tx.Where("ticket_type = ? and status = ? ", ticket.Type, "active").First(&ticket_default_type).Error; err != nil {
		if err := tx.Where("id = ?", 1).First(&ticket_default_type).Error; err != nil {
			tx.Rollback()
			return ticket, errors.New(err.Error())
		}
	}

	if erro := tx.Where("ticket_default_type_id = ? and status = ?", ticket_default_type.ID, "active").First(&ticket_default_timing).Error; erro != nil {
		if err := tx.Where("ticket_default_type_id = ?", 1).First(&ticket_default_timing).Error; err != nil {
			tx.Rollback()
			return ticket, errors.New(err.Error())
		}
	}

	ticket.Priority = ticket_default_timing.TicketPriority

	ticket.Tat = time.Now()
	tat := helpers.GetDuration(ticket_default_timing.Tat)
	ticket.Tat = ticket.ExpiryDate.Add(time.Hour * time.Duration(tat))

	ticket.ExpiryDate = time.Now()
	Duration := helpers.GetDuration(ticket_default_timing.ExpiryDuration)
	ticket.ExpiryDate = ticket.ExpiryDate.Add(time.Hour * time.Duration(Duration))

	ticket.Status = "unresolved"

	stmt := validations.ValidateTicket(ticket)
	if stmt != "validated" {
		return ticket, errors.New(stmt)
	}

	if err := tx.Create(&ticket).Error; err != nil {
		tx.Rollback()
		return ticket, errors.New(err.Error())
	}

	audits.CreateAuditTicket(ticket, db)

	ticket, err = reviewers.CreateTicketReviewer(ticket)
	if err != nil {
		return ticket, err
	}

	// workers.StartTicketClient(ticket.ID,"escalation")
	// workers.StartTicketClient(ticket.ID,"expiration")
	tx.Commit()

	return ticket, err

}
