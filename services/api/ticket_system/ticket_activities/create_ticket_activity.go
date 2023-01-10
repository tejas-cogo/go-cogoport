package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

type TicketActivityService struct {
	TicketActivity models.TicketActivity
}

func CreateTicketActivity(ticket_activity models.TicketActivity) models.TicketActivity {
	db := config.GetDB()
	// result := map[string]interface{}{}

	db.Create(&ticket_activity)
	return ticket_activity
}