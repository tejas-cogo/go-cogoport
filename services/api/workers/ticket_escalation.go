package api

import (
	"encoding/json"
	"time"

	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	helpers "github.com/tejas-cogo/go-cogoport/services/helpers"
)

type TicketWorkerService struct {
	TicketEscalatedPayload models.TicketPayload
}

func TicketEscalation(p models.TicketPayload) error {

	db := config.GetDB()

	var ticket models.Ticket
	var ticket_reviewer models.TicketReviewer
	var ticket_reviewer_new models.TicketReviewer
	var ticket_default_type models.TicketDefaultType
	var ticket_default_timing models.TicketDefaultTiming
	var ticket_default_role models.TicketDefaultRole
	var new_ticket_default_role models.TicketDefaultRole

	tx := db.Begin()

	if err := tx.Where("id = ?", p.TicketID).First(&ticket).Error; err != nil {
		tx.Rollback()
		return err
	}

	if ticket.Status == "unresolved" || ticket.Status == "pending" {

		if err := tx.Where("ticket_type = ?", ticket.Type).First(&ticket_default_type).Error; err != nil {
			tx.Rollback()
			return err
		}

		if err := tx.Where("ticket_default_type_id = ? and status = ?", ticket_default_type.ID, "active").First(&ticket_default_timing).Error; err != nil {
			tx.Rollback()
			return err
		}

		if ticket_default_timing.ID == 0 {
			if err := tx.Where("ticket_default_type_id = ?", 1).First(&ticket_default_timing).Error; err != nil {
				tx.Rollback()
				return err
			}
		}

		ticket.Tat = time.Now()
		Duration := helpers.GetDuration(ticket_default_timing.Tat)
		ticket.ExpiryDate = ticket.ExpiryDate.Add(time.Hour * time.Duration(Duration))

		ticket.ExpiryDate = time.Now()
		Duration = helpers.GetDuration(ticket_default_timing.ExpiryDuration)
		ticket.ExpiryDate = ticket.ExpiryDate.Add(time.Hour * time.Duration(Duration))

		if err := tx.Save(&ticket).Error; err != nil {
			tx.Rollback()
			return err
		}

		if err := tx.Where("ticket_id = ? and status = ?", ticket.ID, "active").First(&ticket_reviewer).Error; err != nil {
			tx.Rollback()
			return err
		}

		ticket_reviewer.Status = "inactive"

		if err := tx.Save(&ticket_reviewer).Error; err != nil {
			tx.Rollback()
			return err
		}

		if err := tx.Where("ticket_default_type_id = ? and status = ? and role_id = ?", ticket_default_type.ID, "active", ticket_reviewer.RoleID).First(&ticket_default_role).Error; err != nil {
			if err := tx.Where("ticket_default_type_id = ? and status = ? and role_id = ?", 1, "active", ticket_reviewer.RoleID).First(&ticket_default_role).Error; err != nil {
				tx.Rollback()
				return err
			}
		}

		if err := tx.Where("ticket_default_type_id = ? and status = ? and role_id = ? and level = ?", ticket_default_type.ID, "active", ticket_reviewer.RoleID, ticket_default_role.Level-1).First(&new_ticket_default_role).Error; err != nil {
			new_ticket_default_role = ticket_default_role
		}

		ticket_reviewer_new.TicketID = ticket.ID
		ticket_reviewer_new.RoleID = new_ticket_default_role.RoleID
		ticket_reviewer_new.UserID = helpers.GetRoleIdUser(new_ticket_default_role.RoleID)

		if err := tx.Create(&ticket_reviewer_new).Error; err != nil {
			tx.Rollback()
			return err
		}

		var ticket_activity models.TicketActivity
		ticket_activity.TicketID = ticket_reviewer_new.TicketID
		ticket_activity.UserID = ticket_reviewer_new.UserID
		ticket_activity.UserType = "system"
		ticket_activity.Type = "Automatically Reviewer Escalated"
		ticket_activity.Status = "escalated"

		if err := tx.Create(&ticket_activity).Error; err != nil {
			tx.Rollback()
			return err
		}

		var ticket_audit models.TicketAudit

		ticket_audit.ObjectId = ticket.ID
		ticket_audit.Action = "escalated"
		ticket_audit.Object = "Ticket"
		data, _ := json.Marshal(ticket)
		ticket_audit.Data = string(data)

		if err := tx.Create(&ticket_audit).Error; err != nil {
			tx.Rollback()
			return err
		}

		tx.Commit()

	}
	return tx.Commit().Error
}
