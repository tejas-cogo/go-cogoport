package api

import (
	"errors"

	"github.com/google/uuid"
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	activities "github.com/tejas-cogo/go-cogoport/services/api/ticket_activities"
	validations "github.com/tejas-cogo/go-cogoport/services/validations"
)

func ReassignTicketReviewer(body models.ReviewerActivity) (models.ReviewerActivity, error) {
	db := config.GetDB()
	tx := db.Begin()
	var err error

	var ticket_reviewer_old models.TicketReviewer
	var ticket_reviewer_active models.TicketReviewer

	if err := tx.Where("ticket_id = ? AND status = 'active'", body.TicketID).First(&ticket_reviewer_active).Error; err != nil {
		tx.Rollback()
		return body, errors.New(err.Error())
	}

	ticket_reviewer_active.Status = "inactive"

	if err := tx.Save(&ticket_reviewer_active).Error; err != nil {
		tx.Rollback()
		return body, errors.New(err.Error())
	}

	if body.ReviewerUserID == uuid.Nil {

		//from ruby client
	} else if body.RoleID == uuid.Nil {
		//from ruby client

	}

	if err := tx.Where("ticket_id = ? AND user_id = ? AND role_id = ?", body.TicketID, body.ReviewerUserID, body.RoleID).Find(&ticket_reviewer_old).Error; err != nil {
		tx.Rollback()
		return body, errors.New(err.Error())
	}

	if ticket_reviewer_old.ID != 0 {
		ticket_reviewer_old.Status = "active"
		ticket_reviewer_old.Level = 3
		if err := tx.Save(&ticket_reviewer_old).Error; err != nil {
			tx.Rollback()
			return body, errors.New(err.Error())
		}

	} else {
		var ticket_reviewer models.TicketReviewer
		ticket_reviewer.TicketID = body.TicketID
		ticket_reviewer.UserID = body.ReviewerUserID
		ticket_reviewer.RoleID = body.RoleID
		ticket_reviewer.UserID = body.ReviewerUserID
		ticket_reviewer.Level = 3

		if body.ReviewerUserID == uuid.Nil {
			var partner_user models.PartnerUser
			tx.Where("role_ids = '{?}'", ticket_reviewer.RoleID).First(&partner_user)
			ticket_reviewer.UserID = partner_user.UserID
			// ticket_reviewer.UserID = helpers.GetRoleIdUser(ticket_reviewer.RoleID)
		}

		stmt := validations.ValidateTicketReviewer(ticket_reviewer)
		if stmt != "validated" {
			return body, errors.New(stmt)
		}
		if err := tx.Create(&ticket_reviewer).Error; err != nil {
			tx.Rollback()
			return body, errors.New(err.Error())
		}

	}

	var filters models.Filter

	filters.TicketActivity.TicketID = body.TicketID
	filters.TicketActivity.UserID = body.PerformedByID
	filters.TicketActivity.UserType = "user"
	filters.TicketActivity.Type = "reviewer_reassigned"
	filters.TicketActivity.Description = body.Description
	filters.TicketActivity.Status = "reassigned"
	activities.CreateTicketActivity(filters)

	tx.Commit()

	return body, err
}
