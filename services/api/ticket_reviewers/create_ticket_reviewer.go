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
			return body, errors.New("Ticket type is invalid!")
		}
	}

	var search_role_user []models.TicketDefaultRole

	if err := txt.Where("ticket_default_type_id = ? and status = ? ", ticket_default_type.ID, "active").Order("level desc").Find(&search_role_user).Error; err != nil {
		if err := txt.Where("ticket_default_type_id = ?", 1).Order("level desc").Find(&search_role_user).Error; err != nil {
			txt.Rollback()
			return body, errors.New("Default Role couldn't be found")
		} else {
			ticket_reviewer = GetNewReviewer(search_role_user, body)
		}
	} else {
		ticket_reviewer = GetNewReviewer(search_role_user, body)
	}

	if ticket_reviewer.UserID == uuid.Nil {
		txt.Rollback()
		return body, errors.New("Reviewer couldn't be found")
	}

	ticket_reviewer.TicketID = body.ID
	ticket_reviewer.ReviewerManagerIDs = helpers.GetUnifiedManagerRmId(ticket_reviewer.UserID)
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

	var ticket_activity models.TicketActivity
	ticket_activity.TicketID = ticket_reviewer.TicketID
	ticket_activity.UserID = ticket_reviewer.UserID
	ticket_activity.UserType = "system"
	ticket_activity.Type = "reviewer_assigned"
	ticket_activity.Status = "assigned"

	stmt3 := validations.ValidateTicketActivity(ticket_activity)
	if stmt3 != "validated" {
		return body, errors.New(stmt)
	}
	if err := txt.Create(&ticket_activity).Error; err != nil {
		txt.Rollback()
		return body, errors.New(err.Error())
	}

	txt.Commit()
	return body, err
}

func GetNewReviewer(search_role_user []models.TicketDefaultRole, body models.Ticket) models.TicketReviewer {

	var ticket_reviewer models.TicketReviewer

	for _, u := range search_role_user {

		if u.UserID != uuid.Nil {
			if u.UserID != body.UserID {
				ticket_reviewer.RoleID = u.RoleID
				ticket_reviewer.UserID = u.UserID
				break
			}
		} else if u.RoleID != uuid.Nil {
			user_id := helpers.GetUnifiedRoleIdUser(u.RoleID, body.UserID.String())
			if user_id != uuid.Nil {
				ticket_reviewer.RoleID = u.RoleID
				ticket_reviewer.UserID = user_id
				break
			}
		}
	}
	return ticket_reviewer
}
