package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func ListTicketAudit() []models.TicketAudit {
	db := config.GetDB()

	var ticket_audit []models.TicketAudit

	result := map[string]interface{}{}
	db.Find(&ticket_audit).Take(&result)

	return ticket_audit
}
