package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func UpdateTicketActivity(body models.TicketActivity) models.TicketActivity {
	db := config.GetDB()
	var ticket_activity models.TicketActivity
	db.Where("id = ?", body.ID).Find(&ticket_activity)

	// if body.Type != "" {
	// 	group.Name = body.Name
	// }
	// if body.Tags != nil {
	// 	group.Tags = body.Tags
	// }
	// if body.UserType != "" {
	// 	ticket_activity.UserType = body.UserType
	// }

	db.Save(&ticket_activity)
	return ticket_activity
}
