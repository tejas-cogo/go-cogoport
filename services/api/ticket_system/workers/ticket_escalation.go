package ticket_system

import (
	"encoding/json"
	"time"

	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	helpers "github.com/tejas-cogo/go-cogoport/services/helpers"
)

type TicketWorkerService struct {
	TicketEscalatedPayload models.TicketEscalatedPayload
}

func TicketEscalation(p models.TicketEscalatedPayload) error {

	db := config.GetDB()

	var ticket models.Ticket
	var ticket_reviewer models.TicketReviewer
	var ticket_reviewer_new models.TicketReviewer
	var ticket_default_timing models.TicketDefaultTiming
	var group_member models.GroupMember
	var group_head models.GroupMember

	tx := db.Begin()

	if err := tx.Where("id = ?", p.TicketID).First(&ticket).Error; err != nil {
		tx.Rollback()
		return err
	}

	if ticket.Status == "unresolved" {

		if err := tx.Where("ticket_type = ?", ticket.Type).First(&ticket_default_timing).Error; err != nil {
			tx.Rollback()
			return err
		}

		if ticket_default_timing.ID == 0 {
			if err := tx.Where("ticket_type = ?", "default").First(&ticket_default_timing).Error; err != nil {
				tx.Rollback()
				return err
			}
		}

		ticket.Tat = ticket_default_timing.Tat

		ticket.ExpiryDate = time.Now()
		Duration := helpers.GetDuration(ticket_default_timing.ExpiryDuration)
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

		if err := tx.Where("ticket_user_id = ? and status = ?", ticket_reviewer.TicketUserID, "active").First(&group_member).Error; err != nil {
			tx.Rollback()
			return err
		}

		group_member.ActiveTicketCount -= 1

		if err := tx.Save(&group_member).Error; err != nil {
			tx.Rollback()
			return err
		}

		if err := tx.Where("group_id = ? and status = ? and hierarchy_level = ?", group_member.GroupID, "active", group_member.HierarchyLevel+1).Order("ActiveTicketCount asc").First(&group_head).Error; err != nil {
			tx.Rollback()
			return err
		}

		ticket_reviewer_new.TicketID = ticket.ID
		ticket_reviewer_new.GroupID = group_head.GroupID
		ticket_reviewer_new.GroupMemberID = group_head.ID
		ticket_reviewer_new.TicketUserID = group_head.TicketUserID

		if err := tx.Create(&ticket_reviewer_new).Error; err != nil {
			tx.Rollback()
			return err
		}

		group_head.ActiveTicketCount += 1

		if err := tx.Save(&group_head).Error; err != nil {
			tx.Rollback()
			return err
		}

		var ticket_activity models.TicketActivity
		ticket_activity.TicketID = ticket_reviewer_new.TicketID
		ticket_activity.TicketUserID = ticket_reviewer_new.TicketUserID
		ticket_activity.UserType = "worker"
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
