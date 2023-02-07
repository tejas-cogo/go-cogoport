package api

import (
	"errors"

	"github.com/google/uuid"
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	audits "github.com/tejas-cogo/go-cogoport/services/api/ticket_audits"
	helpers "github.com/tejas-cogo/go-cogoport/services/helpers"
	validations "github.com/tejas-cogo/go-cogoport/services/validations"
	"gorm.io/gorm"
)

func AllowedUserType() []string {
	return []string{"respond", "rejected", "mark_as_resolved", "reassigned", "escalated", "assigned"}
}

func CreateTicketActivity(body models.Filter) (models.TicketActivity, error) {
	db := config.GetDB()
	var err error
	ticketactivity := body.TicketActivity

	if !(ticketactivity.UserType == "user" || ticketactivity.UserType == "ticket_user") {
		return ticketactivity, errors.New("user type is invalid")
	}

	if body.TicketActivity.Status == "resolved" {
		for _, u := range body.Activity.TicketID {
			ticket_activity := body.TicketActivity
			tx := db.Begin()
			var ticket models.Ticket
			ticket_activity.TicketID = u

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

			if ticket_activity.UserType == "user" {
				// SendTicketActivity(ticket_activity)
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

			if ticket_activity.UserType == "user" {
				// SendTicketActivity(ticket_activity)
			}
			tx.Commit()
		}
	} else if body.TicketActivity.Status == "escalated" {
		for _, u := range body.Activity.TicketID {
			tx := db.Begin()
			ticket_activity := body.TicketActivity
			ticket_activity.TicketID = u
			var ticket_reviewer models.TicketReviewer
			var old_ticket_reviewer models.TicketReviewer
			var ticket_default_type models.TicketDefaultType
			var ticket_default_role models.TicketDefaultRole
			var ticket models.Ticket

			old_ticket_reviewer, err := DeactivateReviewer(u, tx)

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

			if err = tx.Where("ticket_default_type_id = ? and status = ? and level = ?", ticket_default_type.ID, "active", old_ticket_reviewer.Level-1).Order(" level desc").First(&ticket_default_role).Error; err != nil {
				db2 := config.GetCDB().Debug()
				var partner_user models.PartnerUser
				db2.Where("user_id = ? and status = ?", old_ticket_reviewer.UserID, "active").First(&partner_user)
				var manager_user models.PartnerUser

				db2.Where("user_id = ? and status = ?", partner_user.ManagerID, "active").First(&manager_user)
				if manager_user.UserID != uuid.Nil {
					ticket_reviewer.RoleID, _ = uuid.Parse(manager_user.RoleIDs[len(manager_user.RoleIDs)-1])
					ticket_reviewer.UserID = manager_user.UserID
				} else {
					tx.Rollback()
					return ticket_activity, errors.New("cannot escalate further")
				}
			}

			if ticket_reviewer.UserID == uuid.Nil {
				ticket_reviewer.RoleID = ticket_default_role.RoleID
				ticket_reviewer.Level = old_ticket_reviewer.Level - 1
				ticket_reviewer.UserID = helpers.GetRoleIdUser(ticket_default_role.RoleID)
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

			if ticket_activity.UserType == "user" {
				// SendTicketActivity(ticket_activity)
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

func DeactivateReviewer(ID uint, tx *gorm.DB) (models.TicketReviewer, error) {
	var ticket_reviewer models.TicketReviewer
	var err error

	if err := tx.Where("ticket_id = ? and status = ?", ID, "active").First(&ticket_reviewer).Error; err != nil {
		tx.Rollback()
		return ticket_reviewer, errors.New("reviewer not found")
	}

	ticket_reviewer.Status = "inactive"

	if err := tx.Save(&ticket_reviewer).Error; err != nil {
		tx.Rollback()
		return ticket_reviewer, errors.New("cannot update reviewer")
	}

	return ticket_reviewer, err
}
