package api

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	activities "github.com/tejas-cogo/go-cogoport/services/api/ticket_activities"
	audits "github.com/tejas-cogo/go-cogoport/services/api/ticket_audits"
	helpers "github.com/tejas-cogo/go-cogoport/services/helpers"
)

func UpdateTokenTicket(body models.TokenFilter) (models.Ticket, error) {
	db := config.GetDB()
	tx := db

	var err error
	var ticket models.Ticket
	var ticket_token models.TicketToken
	var ticket_reviewer models.TicketReviewer
	var ticket_default_timing models.TicketDefaultTiming
	var ticket_default_role models.TicketDefaultRole
	var ticket_default_type models.TicketDefaultType

	tx.Where("ticket_token = ?", body.TicketToken).First(&ticket_token)
	if ticket_token.Status == "utilized" {
		tx.Rollback()
		return ticket, errors.New("token is already used!")
	}else if ticket_token.TicketID == 0 {
		tx.Rollback()
		return ticket, errors.New("Ticket wasn't created!")
	}

	tx.Where("id = ?", ticket_token.TicketID).First(&ticket)

	if body.Type != "" {
		ticket.Type = body.Type
	}
	if body.Category != "" {
		ticket.Category = body.Category
	}
	if body.Subcategory != "" {
		ticket.Category = body.Category
	}
	if body.NotificationPreferences != nil {
		ticket.NotificationPreferences = body.NotificationPreferences
	}
	if body.Data != nil {
		ticket.Data = body.Data
	}
	if body.Description != "" {
		ticket.Description = body.Description
	}
	if body.IsUrgent != false {
		ticket.IsUrgent = body.IsUrgent
	}

	if err := tx.Where("ticket_type = ? and status = ? ", ticket.Type, "active").First(&ticket_default_type).Error; err != nil {
		tx.Rollback()
		return ticket, errors.New(err.Error())
	}

	if erro := tx.Where("ticket_default_type_id = ? and status = ?", ticket_default_type.ID, "active").First(&ticket_default_timing).Error; erro != nil {
		if err := tx.Where("ticket_default_type_id = ?", 1).First(&ticket_default_timing).Error; err != nil {
			tx.Rollback()
			return ticket, errors.New(err.Error())
		}
	}

	ticket.Priority = ticket_default_timing.TicketPriority

	ticket.Tat = time.Now()
	ticket.ExpiryDate = time.Now()

	Duration := helpers.GetDuration(ticket_default_timing.ExpiryDuration)

	ticket.ExpiryDate = ticket.ExpiryDate.Add(time.Hour * time.Duration(Duration))

	if err := tx.Save(&ticket).Error; err != nil {
		tx.Rollback()
		return ticket, errors.New(err.Error())
	}

	audits.CreateAuditTicket(ticket, db)

	if erro := tx.Where("ticket_default_type_id = ? and status = ?", ticket_default_type.ID, "active").Order("level desc").First(&ticket_default_role).Error; erro != nil {
		if err := tx.Where("ticket_default_type_id = ?", 1).First(&ticket_default_role).Error; err != nil {
			tx.Rollback()
			return ticket, errors.New(err.Error())
		}
	}

	var old_reviewer models.TicketReviewer
	tx.Where("ticket_id = ? and status = ?", ticket.ID, "active").First(&old_reviewer)
	old_reviewer.Status = "inactive"
	tx.Save(&old_reviewer)

	ticket_reviewer.TicketID = ticket.ID
	ticket_reviewer.RoleID = ticket_default_role.RoleID
	ticket_reviewer.UserID = ticket_default_role.UserID
	ticket_reviewer.Status = "active"

	if ticket_reviewer.UserID == uuid.Nil {
		ticket_reviewer.UserID = helpers.GetRoleIdUser(ticket_default_role.RoleID)
	}
	ticket_reviewer.Status = "active"
	ticket_reviewer.Level = ticket_default_role.Level

	var filters models.Filter

	filters.TicketActivity.TicketID = ticket.ID
	filters.TicketActivity.UserType = "system"
	filters.TicketActivity.Type = "reviewer_reassigned"
	filters.TicketActivity.Status = "reassigned"
	activities.CreateTicketActivity(filters)

	ticket_token.Status = "utilized"
	tx.Save(&ticket_reviewer)
	tx.Save(&ticket_token)
	tx.Commit()

	return ticket, err

}
