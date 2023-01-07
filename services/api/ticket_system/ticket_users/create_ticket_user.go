package ticket_system

import (
	"github.com/ChandelShikha/go-cogoport/config"
	"github.com/ChandelShikha/go-cogoport/models"
)

type TicketUserService struct {
	TicketUser models.TicketUser
}

func CreateTicketUser(ticket_user models.TicketUser) models.TicketUser {
	db := config.GetDB()
	// result := map[string]interface{}{}
	db.Create(&ticket_user)
	return ticket_user
}