package ticket_system

import (
	"fmt"
	"github.com/ChandelShikha/go-cogoport/config"
	"github.com/ChandelShikha/go-cogoport/models"
)

func UpdateTicket(id uint, body models.Ticket) models.Ticket {
	db := config.GetDB()
	var ticket models.Ticket
	fmt.Print("Body", body)
	db.Where("id = ?", id).First(&ticket)

	// ticket.Name = body.Name

	db.Save(&ticket)
	return ticket
}