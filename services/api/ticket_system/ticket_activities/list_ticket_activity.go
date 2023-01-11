package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func ListTicketActivity() []models.TicketActivity {
	db := config.GetDB()

	var ticket_activity []models.TicketActivity

	db.Find(&ticket_activity)

	return ticket_activity
}
