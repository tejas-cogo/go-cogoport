package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func UpdateTicketToken(id uint, body models.TicketToken) models.TicketToken {
	db := config.GetDB()
	var ticket_token models.TicketToken
	
	db.Where("id = ?", id).First(&ticket_token)

	ticket_token.TicketUserID = body.TicketUserID

	db.Save(&ticket_token)
	return ticket_token
}