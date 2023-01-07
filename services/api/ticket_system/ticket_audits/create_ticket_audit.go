package ticket_system

import (
	"github.com/ChandelShikha/go-cogoport/config"
	"github.com/ChandelShikha/go-cogoport/models"
)

type TicketAuditService struct {
	TicketAudit models.TicketAudit
}

func CreateTicketAudit(ticket_audit models.TicketAudit) models.TicketAudit {
	db := config.GetDB()
	// result := map[string]interface{}{}
	db.Create(&ticket_audit)
	return ticket_audit
}