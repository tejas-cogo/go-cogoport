package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	audits "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_audits"
)

func UpdateTicket(body models.Ticket) models.Ticket {
	db := config.GetDB()
	var ticket models.Ticket
	db.Where("id = ?", body.ID).First(&ticket)

	ticket.Priority = body.Priority
	

	db.Save(&ticket)

	audits.CreateAuditTicket(ticket, db)
	return ticket
}
