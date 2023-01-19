package ticket_system

import (
	"fmt"

	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	user "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_users"
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

	db.Create(&ticket_activity)
	return ticket_activity
}
