package ticket_system

import (
	"errors"
	"fmt"

	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	activities "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_activities"
)

func ReassignTicketReviewer(body models.ReviewerActivity) (models.ReviewerActivity,error) {
	db := config.GetDB()
	tx := db.Begin()
	var err error

	var ticket_reviewer_old models.TicketReviewer
	var ticket_reviewer_active models.TicketReviewer
	var group_member models.GroupMember

	if err := tx.Where("ticket_id = ? AND status = 'active'", body.TicketID).Find(&ticket_reviewer_active).Error; err != nil {
		tx.Rollback()
		return body, errors.New("Error Occurred!")
	}

	ticket_reviewer_active.Status = "inactive"

	if err := tx.Save(&ticket_reviewer_active).Error; err != nil {
		tx.Rollback()
		return body, errors.New("Error Occurred!")
	}

	if err := tx.Where("id = ?", ticket_reviewer_active.GroupMemberID).First(&group_member).Error; err != nil {
		tx.Rollback()
		return body, errors.New("Error Occurred!")
	}

	group_member.ActiveTicketCount -= 1

	if err := tx.Save(&group_member).Error; err != nil {
		tx.Rollback()
		return body, errors.New("Error Occurred!")
	}

	fmt.Println("edcs", body, "rfd")

	if err := tx.Where("ticket_id = ? AND ticket_user_id = ?", body.TicketID, body.ReviewerUserID).Find(&ticket_reviewer_old).Error; err != nil {
		tx.Rollback()
		return body, errors.New("Error Occurred!")
	}

	if ticket_reviewer_old.ID != 0 {
		ticket_reviewer_old.Status = "active"
		if err := tx.Save(&ticket_reviewer_old).Error; err != nil {
			tx.Rollback()
			return body, errors.New("Error Occurred!")
		}

		if err := tx.Where("id = ?", ticket_reviewer_old.GroupMemberID).First(&group_member).Error; err != nil {
			tx.Rollback()
			return body, errors.New("Error Occurred!")
		}

		group_member.ActiveTicketCount += 1

		if err := tx.Save(&group_member).Error; err != nil {
			tx.Rollback()
			return body, errors.New("Error Occurred!")
		}
	} else {
		var ticket_reviewer models.TicketReviewer
		ticket_reviewer.TicketID = body.TicketID
		ticket_reviewer.TicketUserID = body.ReviewerUserID
		ticket_reviewer.GroupID = body.GroupID
		ticket_reviewer.GroupMemberID = body.GroupMemberID

		if err := tx.Create(&ticket_reviewer).Error; err != nil {
			tx.Rollback()
			return body, errors.New("Error Occurred!")
		}

		if err := tx.Where("id = ?", ticket_reviewer.GroupMemberID).First(&group_member).Error; err != nil {
			tx.Rollback()
			return body, errors.New("Error Occurred!")
		}

		group_member.ActiveTicketCount += 1
		if err := tx.Save(&group_member).Error; err != nil {
			tx.Rollback()
			return body, errors.New("Error Occurred!")
		}
	}

	var filters models.Filter

	filters.TicketActivity.TicketID = body.TicketID
	filters.TicketUser.SystemUserID = body.PerformedByID
	filters.TicketActivity.Type = "reviewer_reassigned"
	filters.TicketActivity.Description = body.Description
	filters.TicketActivity.Status = "reassigned"
	activities.CreateTicketActivity(filters)

	tx.Commit()

	return body, err
}
