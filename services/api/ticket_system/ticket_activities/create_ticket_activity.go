package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	users "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_users"
)

type TicketActivityService struct {
	TicketActivity models.TicketActivity
}

func CreateTicketActivity(ticket_activity models.TicketActivity) models.TicketActivity {
	db := config.GetDB()
	// result := map[string]interface{}{}

	if ticket_activity.UserType == ""{
		var filters models.TicketUser
		filters.ID = ticket_activity.TicketUserID

		ticket_user := users.ListTicketUser(filters)

		for _, u := range ticket_user {
			ticket_activity.UserType = u.Type
			break
		}
	}

	db.Create(&ticket_activity)
	return ticket_activity
}