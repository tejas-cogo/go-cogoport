package api

import (
	"errors"
	_ "time"

	"github.com/google/uuid"
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	audits "github.com/tejas-cogo/go-cogoport/services/api/ticket_audits"

	helpers "github.com/tejas-cogo/go-cogoport/services/helpers"
	validations "github.com/tejas-cogo/go-cogoport/services/validations"
	_ "github.com/tejas-cogo/go-cogoport/tasks"
	_ "github.com/tejas-cogo/go-cogoport/workers"
	"gorm.io/gorm"
)

func CreateTicketActivity(body models.Filter) (string, error) {
	db := config.GetDB()
	var err error
	ticketactivity := body.TicketActivity

	if !(ticketactivity.UserType == "user" || ticketactivity.UserType == "ticket_user" || ticketactivity.UserType == "system") {
		return "", errors.New("user type is invalid")
	}

	if body.TicketActivity.Status == "resolved" {
		tx := db.Begin()
		for _, u := range body.Activity.TicketID {
			ticket_activity := body.TicketActivity

			var ticket models.Ticket
			ticket_activity.TicketID = u

			// validate := validations.ValidateActivityPermission(ticket_activity)
			// if validate == false {
			// 	return ticket_activity, errors.New("You are not authorized to create activity!")
			// }

			if err = tx.Where("id = ? ", u).Where("status = ? or status = ?", "unresolved", "pending").First(&ticket).Error; err != nil {
				tx.Rollback()
				return "", errors.New(err.Error())
			}

			ticket.Status = "closed"

			if err = tx.Save(&ticket).Error; err != nil {
				tx.Rollback()
				return "", errors.New(err.Error())
			}

			DeactivateReviewer(u, tx)

			audits.CreateAuditTicket(ticket, tx)
			stmt := validations.ValidateTicketActivity(ticket_activity)
			if stmt != "validated" {
				return "", errors.New(stmt)
			}

			if err = tx.Create(&ticket_activity).Error; err != nil {

				tx.Rollback()
				return "", errors.New(err.Error())
			}

			if ticket_activity.UserType == "user" {
				// task, err := tasks.ScheduleTicketCommunicationTask(u)
				// if err != nil {
				// 	return ticket_activity, errors.New(err.Error())
				// }
				// Duration := helpers.GetDuration("00h:00m:10s")
				// workers.StartClient((time.Duration(Duration) * time.Minute), task)
			}

		}
		tx.Commit()
		return "Ticket has been resolved!", err
	} else if body.TicketActivity.Status == "requested" {
		tx := db.Begin()
		for _, u := range body.Activity.TicketID {

			ticket_activity := body.TicketActivity

			var ticket models.Ticket
			ticket_activity.TicketID = u

			// validate := validations.ValidateActivityPermission(ticket_activity)
			// if validate == false {
			// 	return ticket_activity, errors.New("You are not authorized to create activity!")
			// }

			if err = tx.Where("id = ?", u).First(&ticket).Error; err != nil {
				tx.Rollback()
				return "", errors.New(err.Error())
			}

			var ticket_default_type models.TicketDefaultType

			if err = tx.Where("id = ? and status = ?", ticket.TicketDefaultTypeID, "active").First(&ticket_default_type).Error; err != nil {
				if err = tx.Where("id = ?", 1).First(&ticket_default_type).Error; err != nil {
					tx.Rollback()
					return "", errors.New(err.Error())
				}
			}

			if len(ticket_default_type.ClosureAuthorizer) != 0 {
				if !helpers.Inslice(ticket_activity.UserID.String(), ticket_default_type.ClosureAuthorizer) {

					ticket.Status = "pending"

					if err = tx.Save(&ticket).Error; err != nil {
						tx.Rollback()
						return "", errors.New(err.Error())
					}

					audits.CreateAuditTicket(ticket, tx)
					stmt := validations.ValidateTicketActivity(ticket_activity)
					if stmt != "validated" {
						return "", errors.New(stmt)
					}
					if err = tx.Create(&ticket_activity).Error; err != nil {
						tx.Rollback()
						return "", errors.New(err.Error())
					}

					if ticket_activity.UserType == "user" {
						// task, err := tasks.ScheduleTicketCommunicationTask(u)
						// if err != nil {
						// 	return ticket_activity, errors.New(err.Error())
						// }
						// Duration := helpers.GetDuration("00h:00m:10s")
						// workers.StartClient((time.Duration(Duration) * time.Minute), task)
					}

				} else {
					var filters models.Filter

					filters.Activity.TicketID = append(filters.Activity.TicketID, u)
					filters.TicketActivity.UserID = ticket_activity.UserID
					filters.TicketActivity.Type = "mark_as_resolved"
					filters.TicketActivity.UserType = ticket_activity.UserType
					filters.TicketActivity.Description = ticket_activity.Description
					filters.TicketActivity.Data = ticket_activity.Data
					filters.TicketActivity.Status = "resolved"

					CreateTicketActivity(filters)
				}
			} else {
				var filters models.Filter

				filters.Activity.TicketID = append(filters.Activity.TicketID, u)
				filters.TicketActivity.UserID = ticket_activity.UserID
				filters.TicketActivity.Type = "mark_as_resolved"
				filters.TicketActivity.UserType = "user"
				filters.TicketActivity.Description = ticket_activity.Description
				filters.TicketActivity.Data = ticket_activity.Data
				filters.TicketActivity.Status = "resolved"

				CreateTicketActivity(filters)
			}
		}
		tx.Commit()
		return "Ticket has been requested!", err
	} else if body.TicketActivity.Status == "rejected" {
		tx := db.Begin()
		for _, u := range body.Activity.TicketID {

			ticket_activity := body.TicketActivity

			var ticket models.Ticket
			ticket_activity.TicketID = u

			// validate := validations.ValidateActivityPermission(ticket_activity)
			// if validate == false {
			// 	return ticket_activity, errors.New("You are not authorized to create activity!")
			// }

			if err = tx.Where("id = ?", u).First(&ticket).Error; err != nil {
				tx.Rollback()
				return "", errors.New(err.Error())
			}
			ticket.Status = "rejected"

			if err = tx.Save(&ticket).Error; err != nil {
				tx.Rollback()
				return "", errors.New(err.Error())
			}

			DeactivateReviewer(u, tx)

			audits.CreateAuditTicket(ticket, tx)
			stmt := validations.ValidateTicketActivity(ticket_activity)
			if stmt != "validated" {
				return "", errors.New(stmt)
			}
			if err = tx.Create(&ticket_activity).Error; err != nil {
				tx.Rollback()
				return "", errors.New(err.Error())
			}

			if ticket_activity.UserType == "user" {
				// task, err := tasks.ScheduleTicketCommunicationTask(u)
				// if err != nil {
				// 	return ticket_activity, errors.New(err.Error())
				// }
				// Duration := helpers.GetDuration("00h:00m:10s")
				// workers.StartClient((time.Duration(Duration) * time.Minute), task)
			}

		}
		tx.Commit()
		return "Ticket has been rejected!", err
	} else if body.TicketActivity.Status == "escalated" {
		tx := db.Begin()
		for _, u := range body.Activity.TicketID {

			ticket_activity := body.TicketActivity

			ticket_activity.TicketID = u
			// validate := validations.ValidateActivityPermission(ticket_activity)
			// if validate == false {
			// 	return ticket_activity, errors.New("You are not authorized to create activity!")
			// }
			var ticket_reviewer models.TicketReviewer
			var old_ticket_reviewer models.TicketReviewer
			var ticket_default_type models.TicketDefaultType
			var ticket_default_role models.TicketDefaultRole
			var ticket models.Ticket

			old_ticket_reviewer, err := DeactivateReviewer(u, tx)
			if old_ticket_reviewer.Level < 1 {
				return "", errors.New("cannot escalate further")
			}
			ticket_reviewer.Level = old_ticket_reviewer.Level - 1
			ticket_reviewer.Status = "active"

			if err != nil {
				tx.Rollback()
				return "", err
			}

			if err = tx.Where("id = ? and status = ?", u, "unresolved").First(&ticket).Error; err != nil {
				tx.Rollback()
				return "", errors.New(err.Error())
			}

			if err = tx.Where("ticket_type = ? and status = ?", ticket.Type, "active").First(&ticket_default_type).Error; err != nil {
				if err = tx.Where("id = ? and status = ?", 1, "active").First(&ticket_default_type).Error; err != nil {
					tx.Rollback()
					return "", errors.New(err.Error())
				}
			}

			if err = tx.Where("ticket_default_type_id = ? and status = ? and level < ?", ticket_default_type.ID, "active", old_ticket_reviewer.Level).Order("level desc").First(&ticket_default_role).Error; err != nil {
				if err = tx.Where("ticket_default_type_id = ? and status = ?", 1, "active").Order("level desc").First(&ticket_default_role).Error; err != nil {
					tx.Rollback()
					return "", errors.New("cannot escalate further")
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

			stmt := validations.ValidateTicketActivity(ticket_activity)
			if stmt != "validated" {
				return "", errors.New(stmt)
			}
			if err = tx.Create(&ticket_reviewer).Error; err != nil {
				tx.Rollback()
				return "", errors.New(err.Error())
			}

			ticket.Status = "escalated"
			audits.CreateAuditTicket(ticket, tx)

			stmt2 := validations.ValidateTicketActivity(ticket_activity)
			if stmt2 != "validated" {
				return "", errors.New(stmt)
			}
			if err = tx.Create(&ticket_activity).Error; err != nil {
				tx.Rollback()
				return "", errors.New(err.Error())
			}
			if ticket_activity.UserType == "user" {
				// task, err := tasks.ScheduleTicketCommunicationTask(u)
				// if err != nil {
				// 	return ticket_activity, errors.New(err.Error())
				// }
				// Duration := helpers.GetDuration("00h:00m:10s")
				// workers.StartClient((time.Duration(Duration) * time.Minute), task)
			}
		}
		tx.Commit()
		return "Ticket has been escalated!", err
	} else if body.TicketActivity.Status == "activity" {
		tx := db.Begin()
		for _, u := range body.Activity.TicketID {

			ticket_activity := body.TicketActivity

			var ticket models.Ticket
			ticket_activity.TicketID = u
			// validate := validations.ValidateActivityPermission(ticket_activity)
			// if validate == false {
			// 	return ticket_activity, errors.New("You are not authorized to create activity!")
			// }

			audits.CreateAuditTicket(ticket, tx)
			stmt := validations.ValidateTicketActivity(ticket_activity)
			if stmt != "validated" {
				return "", errors.New(stmt)
			}
			if err = tx.Create(&ticket_activity).Error; err != nil {
				tx.Rollback()
				return "", errors.New(err.Error())
			}

			if ticket_activity.UserType == "user" {
				// task, err := tasks.ScheduleTicketCommunicationTask(u)
				// if err != nil {
				// 	return ticket_activity, errors.New(err.Error())
				// }
				// Duration := helpers.GetDuration("00h:00m:10s")
				// workers.StartClient((time.Duration(Duration) * time.Minute), task)
			}

		}
		tx.Commit()
		return "Ticket has been activity!", err
	} else if body.TicketActivity.Status == "unresolved" && body.TicketActivity.Type == "resolution_rejected" {
		tx := db.Begin()
		for _, u := range body.Activity.TicketID {

			ticket_activity := body.TicketActivity

			var ticket models.Ticket
			ticket_activity.TicketID = u

			// validate := validations.ValidateActivityPermission(ticket_activity)
			// if validate == false {
			// 	return ticket_activity, errors.New("You are not authorized to create activity!")
			// }

			if err = tx.Where("id = ? and status = ?", u, "pending").First(&ticket).Error; err != nil {
				tx.Rollback()
				return "", errors.New(err.Error())
			}

			ticket.Status = "unresolved"

			if err = tx.Save(&ticket).Error; err != nil {
				tx.Rollback()
				return "", errors.New(err.Error())
			}

			audits.CreateAuditTicket(ticket, tx)
			stmt := validations.ValidateTicketActivity(ticket_activity)
			if stmt != "validated" {
				return "", errors.New(stmt)
			}
			if err = tx.Create(&ticket_activity).Error; err != nil {
				tx.Rollback()
				return "", errors.New(err.Error())
			}

			if ticket_activity.UserType == "user" {
				// task, err := tasks.ScheduleTicketCommunicationTask(u)
				// if err != nil {
				// 	return ticket_activity, errors.New(err.Error())
				// }
				// Duration := helpers.GetDuration("00h:00m:10s")
				// workers.StartClient((time.Duration(Duration) * time.Minute), task)
			}

		}
		tx.Commit()
		return "Pull Requested has been completed!", err
	} else {
		tx := db.Begin()
		var ticket models.Ticket

		ticket_activity := body.TicketActivity
		audits.CreateAuditTicket(ticket, tx)
		stmt := validations.ValidateTicketActivity(ticket_activity)
		if stmt != "validated" {
			return "", errors.New(stmt)
		}
		if err = tx.Create(&ticket_activity).Error; err != nil {
			tx.Rollback()
			return "", errors.New(err.Error())
		}

		if ticket_activity.UserType == "user" {
			// task, err := tasks.ScheduleTicketCommunicationTask(u)
			// if err != nil {
			// 	return ticket_activity, errors.New(err.Error())
			// }
			// Duration := helpers.GetDuration("00h:00m:10s")
			// workers.StartClient((time.Duration(Duration) * time.Minute), task)
		}
		tx.Commit()
		return "Activity is send !", err
	}

}

func DeactivateReviewer(ID uint, tx *gorm.DB) (models.TicketReviewer, error) {
	var ticket_reviewer models.TicketReviewer
	var err error
	var ticket models.Ticket

	if err := tx.Where("ticket_id = ? and status = ?", ID, "active").First(&ticket_reviewer).Error; err != nil {
		tx.Rollback()
		return ticket_reviewer, errors.New("reviewer not found")
	}

	if err := tx.Where("id = ?", ID).First(&ticket).Error; err != nil {
		tx.Rollback()
		return ticket_reviewer, errors.New("reviewer not found")
	}

	if ticket.Status != "unresolved" {
		ticket_reviewer.Status = "closed"
	} else {
		ticket_reviewer.Status = "inactive"
	}

	if err := tx.Save(&ticket_reviewer).Error; err != nil {
		tx.Rollback()
		return ticket_reviewer, errors.New("cannot update reviewer")
	}

	return ticket_reviewer, err
}
