package ticket_system

import (
	"fmt"
	"github.com/ChandelShikha/go-cogoport/config"
	"github.com/ChandelShikha/go-cogoport/models"
)

func UpdateTicketReviewer(id uint, body models.TicketReviewer) models.TicketReviewer {
	db := config.GetDB()
	var ticket_reviewer models.TicketReviewer
	fmt.Print("Body", body)
	db.Where("id = ?", id).First(&ticket_reviewer)

	// ticket_reviewer.Name = body.Name

	db.Save(&ticket_reviewer)
	return ticket_reviewer
}