package api

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	helpers "github.com/tejas-cogo/go-cogoport/services/helpers"
	validations "github.com/tejas-cogo/go-cogoport/services/validations"
)

type TicketReviewerService struct {
	TicketReviewer   models.TicketReviewer
	ReviewerActivity models.ReviewerActivity
}

func CreateTicketReviewer(body models.Ticket) (models.Ticket, error) {
	db := config.GetDB()

	txt := db.Begin()

	var ticket_reviewer models.TicketReviewer
	var ticket_default_type models.TicketDefaultType
	var ticket_default_role models.TicketDefaultRole
	var err error

	if err := txt.Where("ticket_type = ? and status = ?", body.Type, "active").First(&ticket_default_type).Error; err != nil {
		if err := txt.Where("id = ?", 1).First(&ticket_default_type).Error; err != nil {
			txt.Rollback()
			return body, errors.New(err.Error())
		}
	}

	if erro := txt.Where("ticket_default_type_id = ? and status = ?", ticket_default_type.ID, "active").Order("level desc").First(&ticket_default_role).Error; erro != nil {
		if err := txt.Where("ticket_default_type_id = ? ", 1).First(&ticket_default_role).Error; err != nil {
			txt.Rollback()
			return body, errors.New(err.Error())
		}
	}
	ticket_reviewer.TicketID = body.ID
	ticket_reviewer.RoleID = ticket_default_role.RoleID
	ticket_reviewer.UserID = ticket_default_role.UserID

	if ticket_reviewer.UserID == uuid.Nil {
		ticket_reviewer.UserID = helpers.GetRoleIdUser(ticket_reviewer.RoleID)
	}
	// ticket_reviewer.ReviewerManagerIDs = []
	ticket_reviewer.Status = "active"
	ticket_reviewer.Level = ticket_default_role.Level

	stmt := validations.ValidateTicketReviewer(ticket_reviewer)
	fmt.Println("validations", stmt)
	if stmt != "validated" {
		txt.Rollback()

		return body, errors.New(stmt)
	}
	if err := txt.Create(&ticket_reviewer).Error; err != nil {
		txt.Rollback()
		return body, errors.New(err.Error())
	}

	fmt.Println("after", ticket_reviewer)

	var ticket_activity models.TicketActivity
	ticket_activity.TicketID = ticket_reviewer.TicketID
	ticket_activity.UserID = ticket_reviewer.UserID
	ticket_activity.UserType = "system"
	ticket_activity.Type = "reviewer_assigned"
	ticket_activity.Status = "assigned"
	ticket_activity.Status = "assigned"

	stmt3 := validations.ValidateTicketActivity(ticket_activity)
	if stmt3 != "validated" {
		return body, errors.New(stmt)
	}
	if err := txt.Create(&ticket_activity).Error; err != nil {
		txt.Rollback()
		return body, errors.New(err.Error())
	}

	// if ticket_activity.UserType == "internal" {
	// 	// activity.SendTicketActivity(ticket_activity)
	// }

	txt.Commit()
	return body, err
}
