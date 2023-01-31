package ticket_system

import (
	"errors"
	"time"

	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	activities "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_activities"
	audits "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_audits"
	helpers "github.com/tejas-cogo/go-cogoport/services/helpers"
)

func UpdateTokenTicket(body models.TokenFilter) (models.Ticket, error) {
	db := config.GetDB()
	tx := db
	var err error
	var ticket models.Ticket
	var group_member models.GroupMember
	var ticket_token models.TicketToken
	var ticket_reviewer models.TicketReviewer
	var ticket_default_timing models.TicketDefaultTiming
	var ticket_default_group models.TicketDefaultGroup
	var ticket_default_type models.TicketDefaultType

	tx.Where("ticket_token = ?", body.TicketToken).First(&ticket_token)

	tx.Where("id = ?", ticket_token.TicketID).First(&ticket)

	if body.Type != "" {
		ticket.Type = body.Type
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
		return ticket, errors.New("ticket_default_type User Not Found")
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

	if err := tx.Save(&ticket).Error; err != nil {
		tx.Rollback()
		return ticket, errors.New("Ticket couldn't be created")
	}

	audits.CreateAuditTicket(ticket, db)

	if erro := tx.Where("ticket_default_type_id = ? and status = ?", ticket_default_type.ID, "active").First(&ticket_default_group).Error; erro != nil {
		if err := tx.Where("ticket_default_type_id = ?", 1).First(&ticket_default_group).Error; err != nil {
			tx.Rollback()
			return ticket, errors.New("Default Group had issue!")
		}
	}

	var old_reviewer models.TicketReviewer
	tx.Where("ticket_id = ? and status = ?", ticket.ID, "active").First(&old_reviewer)
	old_reviewer.Status = "inactive"
	tx.Save(&old_reviewer)

	tx.Where("id = ? and status = ?", old_reviewer.GroupMemberID, "active").First(&group_member)
	group_member.ActiveTicketCount -= 1
	tx.Save(&group_member)

	ticket_reviewer.TicketID = ticket.ID
	ticket_reviewer.GroupID = ticket_default_group.GroupID

	if ticket_default_group.GroupMemberID > 0 {
		tx.Where("id = ? and status = ?", ticket_default_group.GroupMemberID, "active").First(&group_member)
		ticket_reviewer.GroupMemberID = group_member.ID
	} else {
		tx.Where("group_id = ? and status = ?", ticket_default_group.GroupID, "active").Order("hierarchy_level desc ").Order("active_ticket_count asc").First(&group_member)
		ticket_reviewer.GroupMemberID = group_member.ID
	}
	ticket_reviewer.TicketUserID = group_member.TicketUserID

	tx.Create(&ticket_reviewer)

	group_member.ActiveTicketCount += 1

	tx.Save(&group_member)

	var filters models.Filter

	filters.TicketActivity.TicketID = ticket.ID
	filters.TicketActivity.UserType = "system"
	filters.TicketActivity.Type = "reviewer_reassigned"
	filters.TicketActivity.Status = "reassigned"
	activities.CreateTicketActivity(filters)

	tx.Commit()

	return ticket, err

}
