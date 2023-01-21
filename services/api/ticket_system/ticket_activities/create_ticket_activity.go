package ticket_system

import (
	"fmt"

	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	audits "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_audits"
	user "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_users"
	// tickets "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/tickets"
)

type TicketActivityService struct {
	TicketActivity models.TicketActivity
}

func CreateTicketActivity(body models.Filter) models.TicketActivity {
	db := config.GetDB()
	// result := map[string]interface{}{}

	var ticket_user models.TicketUser

	if body.TicketActivity.UserType == "" {
		if body.TicketActivity.TicketUserID == 0 {
			ticket_user.SystemUserID = body.TicketUser.SystemUserID
		} else {
			ticket_user.ID = body.TicketActivity.TicketUserID
		}

		ticket_user, _ := user.ListTicketUser(ticket_user)
		for _, u := range ticket_user {
			fmt.Println("Fdv", u.ID, "vs")
			body.TicketActivity.UserType = u.Type
			body.TicketActivity.TicketUserID = u.ID
			break
		}
	}
	ticket_activity := body.TicketActivity

	if ticket_activity.Status == "resolved" {
		for _, u := range body.Activity.TicketID {
			var ticket models.Ticket
			db.Model(&ticket).Where("id = ?", u).Update("status", "closed")
			audits.CreateAuditTicket(ticket, db)
			db.Create(&ticket_activity)
		}
	} else if ticket_activity.Status == "rejected" {
		for _, u := range body.Activity.TicketID {
			var ticket models.Ticket
			db.Model(&ticket).Where("id = ?", u).Update("status", "rejected")
			audits.CreateAuditTicket(ticket, db)
			db.Create(&ticket_activity)
		}
	} else if ticket_activity.Status == "escalated" {
		for _, u := range body.Activity.TicketID {
			var group_member models.GroupMember
			var ticket_reviewer models.TicketReviewer
			var ticket models.Ticket

			db.Model(&ticket_reviewer).Where("ticket_id = ? and status = active", u).Update("status", "inactive")

			db.Where("ticket_user_id = ? and status = active", ticket_reviewer.TicketUserID).Find(&group_member)

			ticket_reviewer.TicketID = u
			ticket_reviewer.TicketUserID = group_member.GroupHeadID
			ticket_reviewer.GroupID = group_member.GroupID
			ticket_reviewer.GroupMemberID = group_member.GroupHeadID
			ticket_reviewer.Status = "active"
			db.Create(&ticket_reviewer)
			ticket.Status = "escalated"

			audits.CreateAuditTicket(ticket, db)
			db.Create(&ticket_activity)
		}
	}else if ticket_activity.Status == "activity"{
		db.Create(&ticket_activity)
	} else {
		var ticket models.Ticket
		audits.CreateAuditTicket(ticket, db)
		db.Create(&ticket_activity)
	}
	return ticket_activity
}
