package ticket_system

import (
	"github.com/ChandelShikha/go-cogoport/config"
	"github.com/ChandelShikha/go-cogoport/models"
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