package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	users "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_users"
)

type TicketActivityService struct {
	TicketActivity models.TicketActivity
}

func CreateTicketActivity(body models.Filter) models.TicketActivity {
	db := config.GetDB()
	// result := map[string]interface{}{}

	var ticket_activity models.TicketActivity

	if body.TicketActivity.TicketUserID == 0{
		var filters models.TicketUser
		filters = body.TicketUser
		ticket_user,_ := users.ListTicketUser(filters)

		for _, u := range ticket_user {
			body.TicketActivity.UserType = u.Type
			break
		}
	}

	db.Create(&ticket_activity)
	return ticket_activity
}