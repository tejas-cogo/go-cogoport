package ticket_system

import (
	"fmt"
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func UpdateTicketUser(id uint, body models.TicketUser) models.TicketUser {
	db := config.GetDB()
	var ticket_user models.TicketUser
	fmt.Print("Body", body)
	db.Where("id = ?", id).First(&ticket_user)

	// ticket_user.Name = body.Name

	db.Save(&ticket_user)
	return ticket_user
}