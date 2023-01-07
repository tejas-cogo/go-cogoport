package ticket_system

import (
	"github.com/ChandelShikha/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func ListTicketActivity() []models.TicketActivity {
	db := config.GetDB()

	var ticket_activity []models.TicketActivity

	result := map[string]interface{}{}
	db.Find(&ticket_activity).Take(&result)

	return ticket_activity
}
