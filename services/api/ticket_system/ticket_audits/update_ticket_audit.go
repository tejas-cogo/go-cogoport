package ticket_system

import (
	// "fmt"
	// "github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func UpdateTicketAudit(id uint, body models.TicketAudit) models.TicketAudit {
	// db := config.GetDB()
	var ticket_audit models.TicketAudit
	// fmt.Print("Body", body)
	// db.Where("id = ?", id).First(&ticket_audit)

	// // ticket_audit.Name = body.Name

	// db.Save(&ticket_audit)
	return ticket_audit
}