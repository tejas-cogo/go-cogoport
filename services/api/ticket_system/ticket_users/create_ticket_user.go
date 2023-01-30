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
	var exist_user models.TicketUser

	db.Where("system_user_id = ? and status = ?", ticket_user.SystemUserID, "active").First(&exist_user)

	if exist_user.ID <= 0 {
		db.Create(&ticket_user)
		return ticket_user
	} else {
		return exist_user
	}
	// result := map[string]interface{}{}
}
