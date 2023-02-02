package api

import (
	"errors"
	"fmt"

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
			db := config.GetDB()
			tx := db.Begin()
			var ticket models.Ticket
			ticket_activity.TicketID = u

			if err = tx.Where("id = ?", u).First(&ticket).Error; err != nil {
				tx.Rollback()
				return ticket_activity,  errors.New(err.Error())
			}

			ticket.Status = "closed"

			if err = tx.Save(&ticket).Error; err != nil {
				tx.Rollback()
				return ticket_activity,  errors.New(err.Error())
			}

			DeactivateReviewer(u, tx)

			audits.CreateAuditTicket(ticket, tx)
			stmt := validations.ValidateTicketActivity(ticket_activity)
			if stmt != "validated" {
				return ticket_activity, errors.New(stmt)
			}
			fmt.Println("kljdhbvfksn", ticket_activity)
			if err = tx.Create(&ticket_activity).Error; err != nil {

				tx.Rollback()
				return ticket_activity, errors.New(err.Error())
			}

			if ticket_activity.UserType == "internal" {
				SendTicketActivity(ticket_activity)
			}
			tx.Commit()
		}
	} else if ticket_activity.Status == "rejected" {
		for _, u := range body.Activity.TicketID {
			tx := db.Begin()
			var ticket models.Ticket
			ticket_activity.TicketID = u

			if err = tx.Where("id = ?", u).First(&ticket).Error; err != nil {
				tx.Rollback()
				return ticket_activity,  errors.New(err.Error())
			}
			ticket.Status = "rejected"

			if err = tx.Save(&ticket).Error; err != nil {
				tx.Rollback()
				return ticket_activity,  errors.New(err.Error())
			}

			DeactivateReviewer(u, tx)

			audits.CreateAuditTicket(ticket, tx)
			stmt := validations.ValidateTicketActivity(ticket_activity)
			if stmt != "validated" {
				return ticket_activity, errors.New(stmt)
			}
			if err = tx.Create(&ticket_activity).Error; err != nil {
				tx.Rollback()
				return ticket_activity,  errors.New(err.Error())
			}

			if ticket_activity.UserType == "internal" {
				SendTicketActivity(ticket_activity)
			}
			tx.Commit()
		}
	} else if ticket_activity.Status == "escalated" {
		for _, u := range body.Activity.TicketID {
			tx := db.Begin()
			var group_head models.GroupMember
			ticket_activity.TicketID = u
			var ticket_reviewer models.TicketReviewer
			var ticket models.Ticket

			group_member, err := DeactivateReviewer(u, tx)
			if err != nil {
				tx.Rollback()
				return ticket_activity, err
			}

			if err = tx.Where("group_id = ? and status = ? and hierarchy_level = ?", group_member.GroupID, "active", (group_member.HierarchyLevel)+1).Order("active_ticket_count asc").First(&group_head).Error; err != nil {
				tx.Rollback()
				return ticket_activity,  errors.New(err.Error())
			}

			group_head.ActiveTicketCount += 1

			if err = tx.Save(&group_head).Error; err != nil {
				tx.Rollback()
				return ticket_activity,  errors.New(err.Error())
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
				return ticket_activity,  errors.New(err.Error())
			}

			ticket.Status = "escalated"
			audits.CreateAuditTicket(ticket, tx)

			stmt2 := validations.ValidateTicketActivity(ticket_activity)
			if stmt2 != "validated" {
				return ticket_activity, errors.New(stmt)
			}
			if err = tx.Create(&ticket_activity).Error; err != nil {
				tx.Rollback()
				return ticket_activity, errors.New(err.Error())
			}
			tx.Commit()

		}
	} else if ticket_activity.Status == "activity" {
		for _, u := range body.Activity.TicketID {
			tx := db.Begin()
			var ticket models.Ticket
			ticket_activity.TicketID = u
			audits.CreateAuditTicket(ticket, tx)
			stmt := validations.ValidateTicketActivity(ticket_activity)
			if stmt != "validated" {
				return ticket_activity, errors.New(stmt)
			}
			if err = tx.Create(&ticket_activity).Error; err != nil {
				tx.Rollback()
				return ticket_activity,  errors.New(err.Error())
			}

			if ticket_activity.UserType == "internal" {
				SendTicketActivity(ticket_activity)
			}
			tx.Commit()
		}
	} else {
		var ticket models.Ticket
		tx := db.Begin()
		audits.CreateAuditTicket(ticket, tx)
		stmt := validations.ValidateTicketActivity(ticket_activity)
		if stmt != "validated" {
			return ticket_activity, errors.New(stmt)
		}
		if err = tx.Create(&ticket_activity).Error; err != nil {
			tx.Rollback()
			return ticket_activity,  errors.New(err.Error())
		}
		tx.Commit()
	}

	return ticket_activity, err
}

func DeactivateReviewer(ID uint, tx *gorm.DB) (models.GroupMember, error) {
	var ticket_reviewer models.TicketReviewer
	var group_member models.GroupMember
	var err error

	if err := tx.Where("ticket_id = ? and status = ?", ID, "active").First(&ticket_reviewer).Error; err != nil {
		tx.Rollback()
		return group_member, errors.New(err.Error())
	}

	ticket_reviewer.Status = "inactive"

	if err := tx.Save(&ticket_reviewer).Error; err != nil {
		tx.Rollback()
		return group_member, errors.New(err.Error())
	}

	if err := tx.Where("ticket_user_id = ? and status = ?", ticket_reviewer.TicketUserID, "active").First(&group_member).Error; err != nil {
		tx.Rollback()
		return group_member, errors.New(err.Error())
	}

	group_member.ActiveTicketCount = group_member.ActiveTicketCount - 1

	if err := tx.Save(&group_member).Error; err != nil {
		tx.Rollback()
		return group_member, errors.New(err.Error())
	}

	return group_member, err
}
