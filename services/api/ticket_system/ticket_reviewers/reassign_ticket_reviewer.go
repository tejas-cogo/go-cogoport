package ticket_system

import (
	"fmt"

	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	activities "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_activities"
	// activities "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_activities"
)

func ReassignTicketReviewer(body models.ReviewerActivity) string {
	db := config.GetDB()
	// result := map[string]interface{}{}

	var ticket_reviewer_old models.TicketReviewer
	var ticket_reviewer_active models.TicketReviewer
	var group_member models.GroupMember

	db.Where("ticket_id = ? AND status = 'active'", body.TicketID).Find(&ticket_reviewer_active)

	ticket_reviewer_active.Status = "inactive"
	db.Save(&ticket_reviewer_active)

	db.Where("id = ?", ticket_reviewer_active.GroupMemberID).First(&group_member)
	group_member.ActiveTicketCount -= 1
	db.Save(&group_member)

	fmt.Println("edcs", body, "rfd")

	db.Where("ticket_id = ? AND ticket_user_id = ?", body.TicketID, body.ReviewerUserID).Find(&ticket_reviewer_old)

	if ticket_reviewer_old.ID != 0 {
		ticket_reviewer_old.Status = "active"
		db.Save(&ticket_reviewer_old)

		db.Where("id = ?", ticket_reviewer_old.GroupMemberID).First(&group_member)
		group_member.ActiveTicketCount += 1
		db.Save(&group_member)
	} else {
		var ticket_reviewer models.TicketReviewer
		ticket_reviewer.TicketID = body.TicketID
		ticket_reviewer.TicketUserID = body.ReviewerUserID
		ticket_reviewer.GroupID = body.GroupID
		ticket_reviewer.GroupMemberID = body.GroupMemberID
		db.Create(&ticket_reviewer)

		db.Where("id = ?", ticket_reviewer.GroupMemberID).First(&group_member)
		group_member.ActiveTicketCount += 1
		db.Save(&group_member)
	}

	var filters models.Filter

	filters.TicketActivity.TicketID = body.TicketID
	filters.TicketUser.SystemUserID = body.PerformedByID
	filters.TicketActivity.Type = "reviewer_reassigned"
	filters.TicketActivity.Description = body.Description
	filters.TicketActivity.Status = "reassigned"
	activities.CreateTicketActivity(filters)

	return "Reassigned Successfully"
}
