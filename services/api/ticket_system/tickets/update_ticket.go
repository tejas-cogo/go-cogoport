package ticket_system

import (
	"errors"

	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	audits "github.com/tejas-cogo/go-cogoport/services/api/ticket_system/ticket_audits"
)

func UpdateTicket(body models.Ticket) (models.Ticket,error) {
	db := config.GetDB()
	tx := db.Begin()
	var err error

	var ticket models.Ticket

	if err := tx.Where("id = ?", body.ID).First(&ticket).Error; err != nil {
		tx.Rollback()
		return body, errors.New("Error Occurred!")
	}

	if body.Priority != ticket.Priority {
		ticket.Priority = body.Priority
	}

	if err := tx.Save(&ticket).Error; err != nil {
		tx.Rollback()
		return body, errors.New("Error Occurred!")
	}

	audits.CreateAuditTicket(ticket, db)
	
	tx.Commit()
	return ticket,err
}
