package ticket_system

import (
	"github.com/ChandelShikha/go-cogoport/config"
	"github.com/ChandelShikha/go-cogoport/models"
)

type TicketDefaultGroupService struct {
	TicketDefaultGroup models.TicketDefaultGroup
}

func CreateTicketDefaultGroup(ticket_default_group models.TicketDefaultGroup) models.TicketDefaultGroup {
	db := config.GetDB()
	// result := map[string]interface{}{}
	db.Create(&ticket_default_group)
	return ticket_default_group
}