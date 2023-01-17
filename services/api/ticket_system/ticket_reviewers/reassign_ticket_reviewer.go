package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	activities "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_activities"
)

func ReassignTicketReviewer(activity models.Activity, body models.TicketReviewer) models.TicketReviewer {
	db := config.GetDB()
	// result := map[string]interface{}{}

	var ticket_reviewer_old []models.TicketReviewer
	var ticket_reviewer_active models.TicketReviewer
	var ticket_reviewer_new models.TicketReviewer

	ticket_reviewer_new.TicketID = body.TicketID
	ticket_reviewer_new.GroupID = body.GroupID
	ticket_reviewer_new.GroupMemberID = body.GroupMemberID

	db.Where("ticket_id = ?", body.TicketID)
	db.Find(&ticket_reviewer_old)

	for _, u := range ticket_reviewer_old {
		if u == ticket_reviewer_new {
			u.Status = "active"
			db.Save(&u)
		}
	}

	db.Where("id = ?", body.ID)
	db.Find(&ticket_reviewer_active)

	ticket_reviewer_active.Status = "inactive"
	db.Save(&ticket_reviewer_active)

	db.Create(&ticket_reviewer_new)

	var filters models.Filter

	filters.TicketActivity.TicketID = body.TicketID
	filters.TicketUser.SystemUserID = activity.SystemUserID
	filters.TicketActivity.Type = "Reviewer Reassigned"
	filters.TicketActivity.Data = activity.Data

	activities.CreateTicketActivity(filters)

	return ticket_reviewer_new
}
