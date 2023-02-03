package api

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
	var ticket_default_type models.TicketDefaultType
	var ticket_default_timing models.TicketDefaultTiming
	var group_member models.GroupMember
	var group_head models.GroupMember

	tx := db.Begin()

	if err := tx.Where("id = ?", p.TicketID).First(&ticket).Error; err != nil {
		tx.Rollback()
		return err
	}

	if ticket.Status == "unresolved" {

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

		var ticket_default_group models.TicketDefaultGroup
		if err := tx.Where("group_id = ? and status = ?", group_member.GroupID, "active").First(&ticket_default_group).Error; err != nil {
			tx.Rollback()
			return err
		}

		var escalated_ticket_default_group models.TicketDefaultGroup

		if err := tx.Where("ticket_default_type_id = ? and status = ? and level = ?", ticket_default_group.TicketDefaultTypeID, "active", ticket_default_group.Level-1).First(&escalated_ticket_default_group).Error; err != nil {
			if err := tx.Where("ticket_default_type_id = ? and status = ? and level = ?", ticket_default_group.TicketDefaultTypeID, "active", ticket_default_group.Level).First(&escalated_ticket_default_group).Error; err != nil {
				tx.Rollback()
				return err
			}
		}

		if escalated_ticket_default_group.GroupMemberID == 0 {
			if ticket_default_group.Level != escalated_ticket_default_group.Level {
				if err := tx.Where("group_id = ?  and status = ?", escalated_ticket_default_group.GroupID, "active").Order("hierarchy_level desc").Order("active_ticket_count asc").First(&group_head).Error; err != nil {
					tx.Rollback()
					return err
				}
			} else {
				if err := tx.Where("group_id = ?  and status = ? and hierarchy_level = ?", escalated_ticket_default_group.GroupID, "active", (group_member.HierarchyLevel)-1).Order("active_ticket_count asc").First(&group_head).Error; err != nil {
					tx.Rollback()
					return err
				}
			}
		} else {
			if ticket_default_group.Level != escalated_ticket_default_group.Level {
				if err := tx.Where("id = ?  and status = ?", escalated_ticket_default_group.GroupMemberID, "active").First(&group_head).Error; err != nil {
					tx.Rollback()
					return err
				}
			}

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
