package api

import (
	"encoding/json"

	_ "github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	"gorm.io/gorm"
)

type TicketAuditService struct {
	TicketAudit models.TicketAudit
}

func CreateAuditTicket(ticket models.Ticket, db *gorm.DB) string {
	var ticket_audit models.TicketAudit

	ticket_audit.ObjectId = ticket.ID
	if ticket.Status == "unresolved" {
		ticket_audit.Action = "created"
	} else if ticket.Status == "inactive" {
		ticket_audit.Action = "deleted"
	} else {
		ticket_audit.Action = "updated"
	}
	ticket_audit.Object = "ticket"
	data, _ := json.Marshal(ticket)
	ticket_audit.Data = string(data)

	db.Create(&ticket_audit)

	return "Audit created successfully"

}
