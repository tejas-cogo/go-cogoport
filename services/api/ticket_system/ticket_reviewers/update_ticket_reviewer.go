package ticket_system

import (
	"fmt"
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func UpdateTicketReviewer(body models.TicketReviewer) models.TicketReviewer {
	db := config.GetDB()
	var ticket_reviewer models.TicketReviewer
	var reviewer_old models.TicketReviewer
	fmt.Print("Body", body)
	db.Where("ticket_user_id = ?", body.TicketUserID)
	db.Where("ticket_id = ?", body.TicketID)
	db.First(&ticket_reviewer)

	if ticket_reviewer == reviewer_old{
		db.Create(&ticket_reviewer)	
	}

	db.Save(&ticket_reviewer)
	return ticket_reviewer
}