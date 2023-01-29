package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func UpdateTicketDefaultGroup(body models.TicketDefaultGroup) models.TicketDefaultGroup {
	db := config.GetDB()
	var ticket_default_group models.TicketDefaultGroup
	db.Where("id = ?", body.ID).Find(&ticket_default_group)

	if body.TicketDefaultTypeID > 0 {
		ticket_default_group.TicketDefaultTypeID = body.TicketDefaultTypeID
	}
	if body.GroupID != 0 {
		ticket_default_group.GroupID = body.GroupID
	}
	if body.Status != "" {
		ticket_default_group.Status = body.Status
	}

	db.Save(&ticket_default_group)
	return ticket_default_group
}
