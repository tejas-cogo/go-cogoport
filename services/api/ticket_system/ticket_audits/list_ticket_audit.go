package ticket_system

import (
	"github.com/ChandelShikha/go-cogoport/config"
	"github.com/ChandelShikha/go-cogoport/models"
)

func ListTicketAudit() []models.TicketAudit {
	db := config.GetDB()

	var ticket_audit []models.TicketAudit

	result := map[string]interface{}{}
	db.Find(&ticket_audit).Take(&result)

	return ticket_audit
}
