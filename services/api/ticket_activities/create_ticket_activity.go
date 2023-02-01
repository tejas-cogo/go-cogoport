package api

import (
	"errors"

	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	audits "github.com/tejas-cogo/go-cogoport/services/api/ticket_audits"
	user "github.com/tejas-cogo/go-cogoport/services/api/ticket_users"
	validations "github.com/tejas-cogo/go-cogoport/services/validations"
	"gorm.io/gorm"
)

type TicketActivityService struct {
	TicketActivity models.TicketActivity
}

func CreateTicketActivity(body models.Filter) (models.TicketActivity, error) {
	db := config.GetDB()
	tx := db.Begin()
	var err error

	var ticket_user models.TicketUserFilter

	if body.TicketActivity.UserType == "" {
		if body.TicketActivity.TicketUserID == 0 {
			ticket_user.SystemUserID = body.TicketUserFilter.SystemUserID
		} else {
			ticket_user.ID = body.TicketUserFilter.ID
		}

		ticket_user, _, _ := user.ListTicketUser(ticket_user)
		for _, u := range ticket_user {
			body.TicketActivity.UserType = u.Type
			body.TicketActivity.TicketUserID = u.ID
			break
		}
	}
	ticket_activity := body.TicketActivity

	if ticket_activity.Status == "resolved" {
		for _, u := range body.Activity.TicketID {
			var ticket models.Ticket
			ticket_activity.TicketID = u

			if err = tx.Where("id = ?", u).First(&ticket).Error; err != nil {
				tx.Rollback()
				return ticket_activity, errors.New("Ticket not found")
			}

			ticket.Status = "closed"

			if err = tx.Save(&ticket).Error; err != nil {
				tx.Rollback()
				return ticket_activity, errors.New("Ticket couldn't be saved")
			}

			DeactivateReviewer(u, tx)

			audits.CreateAuditTicket(ticket, tx)
			stmt := validations.ValidateTicketActivity(ticket_activity)
			if stmt != "validated" {
				return ticket_activity, errors.New(stmt)
			}
			if err = tx.Create(&ticket_activity).Error; err != nil {
				tx.Rollback()
				return ticket_activity, errors.New("Activity couldn't be created")
			}

			if ticket_activity.UserType == "internal" {
				SendTicketActivity(ticket_activity)
			}
		}
	} else if ticket_activity.Status == "rejected" {
		for _, u := range body.Activity.TicketID {
			var ticket models.Ticket
			ticket_activity.TicketID = u

			if err = tx.Where("id = ?", u).First(&ticket).Error; err != nil {
				tx.Rollback()
				return ticket_activity, errors.New("ticket couldn't be find")
			}
			ticket.Status = "rejected"

			if err = tx.Save(&ticket).Error; err != nil {
				tx.Rollback()
				return ticket_activity, errors.New("Ticket couldn't be saved")
			}

			DeactivateReviewer(u, tx)

			audits.CreateAuditTicket(ticket, tx)
			stmt := validations.ValidateTicketActivity(ticket_activity)
			if stmt != "validated" {
				return ticket_activity, errors.New(stmt)
			}
			if err = tx.Create(&ticket_activity).Error; err != nil {
				tx.Rollback()
				return ticket_activity, errors.New("Activity couldn't be created")
			}

			if ticket_activity.UserType == "internal" {
				SendTicketActivity(ticket_activity)
			}
		}
	} else if ticket_activity.Status == "escalated" {
		for _, u := range body.Activity.TicketID {
			var group_head models.GroupMember
			ticket_activity.TicketID = u
			var ticket_reviewer models.TicketReviewer
			var ticket models.Ticket

			group_member, err := DeactivateReviewer(u, tx)
			if err != nil {
				tx.Rollback()
				return ticket_activity, errors.New("Escalated Group Member not found")
			}

			if err = tx.Where("group_id = ? and status = ? and hierarchy_level = ?", group_member.GroupID, "active", (group_member.HierarchyLevel)+1).Order("active_ticket_count asc").First(&group_head).Error; err != nil {
				tx.Rollback()
				return ticket_activity, errors.New("Escalated Group Member not found")
			}

			group_head.ActiveTicketCount += 1

			if err = tx.Save(&group_head).Error; err != nil {
				tx.Rollback()
				return ticket_activity, errors.New("Escalated Group Member couldn't be saved")
			}

			ticket_reviewer.TicketID = u
			ticket_reviewer.TicketUserID = group_head.TicketUserID
			ticket_reviewer.GroupID = group_head.GroupID
			ticket_reviewer.GroupMemberID = group_head.ID
			ticket_reviewer.Status = "active"

			stmt := validations.ValidateTicketActivity(ticket_activity)
			if stmt != "validated" {
				return ticket_activity, errors.New(stmt)
			}
			if err = tx.Create(&ticket_reviewer).Error; err != nil {
				tx.Rollback()
				return ticket_activity, errors.New("Reviewer couldn't be created")
			}

			ticket.Status = "escalated"
			audits.CreateAuditTicket(ticket, tx)

			stmt2 := validations.ValidateTicketActivity(ticket_activity)
			if stmt2 != "validated" {
				return ticket_activity, errors.New(stmt)
			}
			if err = tx.Create(&ticket_activity).Error; err != nil {
				tx.Rollback()
				return ticket_activity, errors.New("Activity couldn't be created")
			}

		}
	} else if ticket_activity.Status == "activity" {
		for _, u := range body.Activity.TicketID {
			var ticket models.Ticket
			ticket_activity.TicketID = u
			audits.CreateAuditTicket(ticket, tx)
			stmt := validations.ValidateTicketActivity(ticket_activity)
			if stmt != "validated" {
				return ticket_activity, errors.New(stmt)
			}
			if err = tx.Create(&ticket_activity).Error; err != nil {
				tx.Rollback()
				return ticket_activity, errors.New("Activity couldn't be created")
			}

			if ticket_activity.UserType == "internal" {
				SendTicketActivity(ticket_activity)
			}
		}
	} else {
		var ticket models.Ticket
		audits.CreateAuditTicket(ticket, tx)
		stmt := validations.ValidateTicketActivity(ticket_activity)
		if stmt != "validated" {
			return ticket_activity, errors.New(stmt)
		}
		if err = tx.Create(&ticket_activity).Error; err != nil {
			tx.Rollback()
			return ticket_activity, errors.New("Activity couldn't be created")
		}
	}

	tx.Commit()

	return ticket_activity, err
}

func DeactivateReviewer(ID uint, tx *gorm.DB) (models.GroupMember, error) {
	var ticket_reviewer models.TicketReviewer
	var group_member models.GroupMember
	var err error

	if err := tx.Where("ticket_id = ? and status = ?", ID, "active").First(&ticket_reviewer).Error; err != nil {
		tx.Rollback()
		return group_member, err
	}

	ticket_reviewer.Status = "inactive"

	if err := tx.Save(&ticket_reviewer).Error; err != nil {
		tx.Rollback()
		return group_member, err
	}

	if err := tx.Where("ticket_user_id = ? and status = ?", ticket_reviewer.TicketUserID, "active").First(&group_member).Error; err != nil {
		tx.Rollback()
		return group_member, err
	}

	group_member.ActiveTicketCount = group_member.ActiveTicketCount - 1

	if err := tx.Save(&group_member).Error; err != nil {
		tx.Rollback()
		return group_member, err
	}

	return group_member, err
}
