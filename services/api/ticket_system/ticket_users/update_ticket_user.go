package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func UpdateTicketUser(body models.TicketUser) models.TicketUser {
	db := config.GetDB()
	var ticket_user models.TicketUser
	db.Where("id = ?", body.ID).First(&ticket_user)

	if body.Type != ticket_user.Type {
		ticket_user.Type = body.Type
	}
	if body.RoleID != ticket_user.RoleID {
		ticket_user.RoleID = body.RoleID
	}
	if body.Source != ticket_user.Source {
		ticket_user.Source = body.Source
	}

	db.Save(&ticket_user)
	return ticket_user
}
