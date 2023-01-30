package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	audits "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_audits"
)

func UpdateTicket(body models.Ticket) (string,error,models.Ticket) {
	db := config.GetDB()
	tx := db.Begin()
	var err error

	var ticket models.Ticket

	if err := tx.Where("id = ?", body.ID).First(&ticket).Error; err != nil {
		tx.Rollback()
		return "Error Occurred!", err, body
	}

	if body.Priority != ticket.Priority {
		ticket.Priority = body.Priority
	}

	if err := tx.Save(&ticket).Error; err != nil {
		tx.Rollback()
		return "Error Occurred!", err, body
	}

	audits.CreateAuditTicket(ticket, db)
	
	tx.Commit()
	return "Successfully Deleted!", err, ticket
}
