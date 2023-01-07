package ticket_system

import (
	"fmt"
	"github.com/ChandelShikha/go-cogoport/config"
	"github.com/ChandelShikha/go-cogoport/models"
)

func UpdateTicketActivity(id uint, body models.TicketActivity) models.TicketActivity {
	db := config.GetDB()
	var ticket_activity models.TicketActivity
	fmt.Print("Body", body)
	db.Where("id = ?", id).First(&ticket_activity)

	// ticket_activity.Name = body.Name

	db.Save(&ticket_activity)
	return ticket_activity
}