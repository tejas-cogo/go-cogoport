package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func ListTicketUser() []models.TicketUser {
	db := config.GetDB()

	var ticket_user []models.TicketUser

	result := map[string]interface{}{}
	db.Find(&ticket_user).Take(&result)

	return ticket_user
}
