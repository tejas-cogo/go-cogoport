package ticket_system

import (
	"github.com/ChandelShikha/go-cogoport/config"
	"github.com/ChandelShikha/go-cogoport/models"
)

func DeleteTicketAudit(id uint) uint{
	db := config.GetDB()

	var ticket_audit models.TicketAudit

	db.Model(&ticket_audit).Where("id = ?", id).Update("status","inactive")

	db.Where("id = ?", id).Delete(&ticket_audit)

	return id
}