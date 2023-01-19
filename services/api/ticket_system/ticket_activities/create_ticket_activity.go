package ticket_system

import (
	"encoding/json"
	"fmt"

	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	user "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_users"
	"gorm.io/gorm"
	// tickets "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/tickets"
)

type TicketActivityService struct {
	TicketActivity models.TicketActivity
}

func CreateAuditTicket(ticket models.Ticket, db *gorm.DB) int {
	var ticket_audit models.TicketAudit

	ticket_audit.ObjectId = ticket.ID
	ticket_audit.Action = ticket.Status
	ticket_audit.Object = "ticket"
	data, _ := json.Marshal(ticket)
	ticket_audit.Data = string(data)

	db.Create(&ticket_audit)

	return 0
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
		var ticket models.Ticket

		db.Model(&ticket).Where("id = ?", body.TicketActivity.TicketID).Update("status", "closed")

		CreateAuditTicket(ticket, db)

	} else if ticket_activity.Status == "rejected" {
		var ticket models.Ticket

		db.Model(&ticket).Where("id = ?", body.TicketActivity.TicketID).Update("status", "rejected")

		CreateAuditTicket(ticket, db)
	}

	db.Create(&ticket_activity)

	return ticket_activity
}
