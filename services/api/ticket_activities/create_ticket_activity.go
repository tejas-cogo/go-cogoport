package api

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	audits "github.com/tejas-cogo/go-cogoport/services/api/ticket_audits"
	helpers "github.com/tejas-cogo/go-cogoport/services/helpers"
	validations "github.com/tejas-cogo/go-cogoport/services/validations"
	"gorm.io/gorm"
)

type TicketActivityService struct {
	ticket_activity models.TicketActivity
}

func CreateTicketActivity(body models.Filter) (models.TicketActivity, error) {
	db := config.GetDB()
	var err error
	var ticket_user models.TicketUser

	//reviewer assigned
	if body.TicketActivity.UserType == "system" {
		db.Where("system_user_id = ? ", body.Activity.PerformedByID).Find(&ticket_user)

		body.TicketActivity.UserID = body.Activity.PerformedByID

	} else if body.TicketActivity.Status == "resolved" || body.TicketActivity.Status == "rejected" || body.TicketActivity.Status == "escalated" || body.TicketActivity.Status == "reviewer_reassigned" {

	}

	if body.TicketActivity.Status == "resolved" {
		for _, u := range body.Activity.TicketID {

			ticket_activity := body.TicketActivity
			tx := db.Begin()
			var ticket models.Ticket
			ticket_activity.TicketID = u

			fmt.Println("ticket_activity", ticket_activity.ID)

			if err = tx.Where("id = ?", u).First(&ticket).Error; err != nil {
				tx.Rollback()
				return ticket_activity, errors.New(err.Error())
			}

			ticket.Status = "closed"

			if err = tx.Save(&ticket).Error; err != nil {
				tx.Rollback()
				return ticket_activity, errors.New(err.Error())
			}

			DeactivateReviewer(u, tx)

			audits.CreateAuditTicket(ticket, tx)
			stmt := validations.ValidateTicketActivity(ticket_activity)
			if stmt != "validated" {
				return ticket_activity, errors.New(stmt)
			}

			if err = tx.Create(&ticket_activity).Error; err != nil {

				tx.Rollback()
				return ticket_activity, errors.New(err.Error())
			}

			if ticket_activity.UserType == "internal" {
				SendTicketActivity(ticket_activity)
			}
			tx.Commit()
		}
	} else if body.TicketActivity.Status == "rejected" {
		for _, u := range body.Activity.TicketID {
			tx := db.Begin()
			ticket_activity := body.TicketActivity
			var ticket models.Ticket
			ticket_activity.TicketID = u

			if err = tx.Where("id = ?", u).First(&ticket).Error; err != nil {
				tx.Rollback()
				return ticket_activity, errors.New(err.Error())
			}
			ticket.Status = "rejected"

			if err = tx.Save(&ticket).Error; err != nil {
				tx.Rollback()
				return ticket_activity, errors.New(err.Error())
			}

			DeactivateReviewer(u, tx)

			audits.CreateAuditTicket(ticket, tx)
			stmt := validations.ValidateTicketActivity(ticket_activity)
			if stmt != "validated" {
				return ticket_activity, errors.New(stmt)
			}
			if err = tx.Create(&ticket_activity).Error; err != nil {
				tx.Rollback()
				return ticket_activity, errors.New(err.Error())
			}

			if ticket_activity.UserType == "internal" {
				SendTicketActivity(ticket_activity)
			}
			tx.Commit()
		}
	} else if body.TicketActivity.Status == "escalated" {
		for _, u := range body.Activity.TicketID {
			tx := db.Begin()
			ticket_activity := body.TicketActivity
			ticket_activity.TicketID = u
			var ticket_reviewer models.TicketReviewer
			var ticket_default_type models.TicketDefaultType
			var ticket models.Ticket

			ticket_default_role, err := DeactivateReviewer(u, tx)
			if err != nil {
				tx.Rollback()
				return ticket_activity, err
			}

			if err = tx.Where("id = ? and status = ?", u, "unresolved").First(&ticket).Error; err != nil {
				tx.Rollback()
				return ticket_activity, errors.New(err.Error())
			}

			if err = tx.Where("ticket_type = ? and status = ?", ticket.Type, "active").First(&ticket_default_type).Error; err != nil {
				if err = tx.Where("id = ? and status = ?", 1, "active").First(&ticket_default_type).Error; err != nil {
					tx.Rollback()
					return ticket_activity, errors.New(err.Error())
				}
			}

			if err = tx.Where("ticket_default_type_id = ? and status = ? and level<?", ticket_default_type.ID, "active", ticket_default_role.Level).Order(" level desc").First(&ticket_default_role).Error; err != nil {
				if err = tx.Where("ticket_default_type_id = ? and status = ? and level<?", 1, "active", ticket_default_role.Level).Order(" level desc").First(&ticket_default_role).Error; err != nil {
					tx.Rollback()
					return ticket_activity, errors.New(err.Error())
				}
			}

			if ticket_default_role.UserID == uuid.Nil {

				ticket_reviewer.RoleID = ticket_default_role.RoleID
				ticket_reviewer.UserID = helpers.GetRoleIdUser(ticket_default_role.RoleID)
			} else {
				ticket_reviewer.RoleID = ticket_default_role.RoleID
				ticket_reviewer.UserID = ticket_default_role.UserID
			}

			ticket_reviewer.TicketID = u
			ticket_reviewer.Status = "active"

			stmt := validations.ValidateTicketActivity(ticket_activity)
			if stmt != "validated" {
				return ticket_activity, errors.New(stmt)
			}
			if err = tx.Create(&ticket_reviewer).Error; err != nil {
				tx.Rollback()
				return ticket_activity, errors.New(err.Error())
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
	} else if body.TicketActivity.Status == "activity" {
		for _, u := range body.Activity.TicketID {
			ticket_activity := body.TicketActivity
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
				return ticket_activity, errors.New(err.Error())
			}

			if ticket_activity.UserType == "internal" {
				SendTicketActivity(ticket_activity)
			}
			tx.Commit()
		}
	} else {
		var ticket models.Ticket
		tx := db.Begin()
		ticket_activity := body.TicketActivity
		audits.CreateAuditTicket(ticket, tx)
		stmt := validations.ValidateTicketActivity(ticket_activity)
		if stmt != "validated" {
			return ticket_activity, errors.New(stmt)
		}
		if err = tx.Create(&ticket_activity).Error; err != nil {
			tx.Rollback()
			return ticket_activity, errors.New(err.Error())
		}
		tx.Commit()
	}

	return body.TicketActivity, err
}

func DeactivateReviewer(ID uint, tx *gorm.DB) (models.TicketDefaultRole, error) {
	var ticket_reviewer models.TicketReviewer
	var ticket_default_role models.TicketDefaultRole
	var err error

	if err := tx.Where("ticket_id = ? and status = ?", ID, "active").First(&ticket_reviewer).Error; err != nil {
		tx.Rollback()
		return ticket_default_role, errors.New(err.Error())
	}

	ticket_reviewer.Status = "inactive"

	if err := tx.Save(&ticket_reviewer).Error; err != nil {
		tx.Rollback()
		return ticket_default_role, errors.New(err.Error())
	}

	if err := tx.Where("user_id = ? and status = ? and role_id = ?", ticket_reviewer.UserID, "active", ticket_reviewer.RoleID).First(&ticket_default_role).Error; err != nil {
		tx.Rollback()
		return ticket_default_role, errors.New(err.Error())
	}

	return ticket_default_role, err
}
