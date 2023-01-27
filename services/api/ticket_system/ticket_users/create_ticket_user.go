package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

type TicketUserService struct {
	TicketUser models.TicketUser
}

func CreateTicketUser(ticket_user models.TicketUser) models.TicketUser {
	db := config.GetDB()
	ticket_user.Status = "active"
	db.Create(&ticket_user)

	return ticket_user
}
