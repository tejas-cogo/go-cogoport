package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	audits "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_audits"
)

func UpdateTicket(body models.Ticket) (models.Ticket, error) {
	db := config.GetDB()
	var ticket models.Ticket
	var err error
	tx := db.Begin()
	if err := tx.Where("id = ?", body.ID).First(&ticket).Error; err != nil {
		tx.Rollback()
		return ticket, err
	}

	ticket.Priority = body.Priority

	if err := tx.Save(&ticket).Error; err != nil {
		tx.Rollback()
		return ticket, err
	}

	audits.CreateAuditTicket(ticket, db)
	return ticket, err
}
