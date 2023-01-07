package ticket_system

import (
	"fmt"
	"github.com/ChandelShikha/go-cogoport/config"
	"github.com/ChandelShikha/go-cogoport/models"
)

func UpdateTicketDefaultGroup(id uint, body models.TicketDefaultGroup) models.TicketDefaultGroup {
	db := config.GetDB()
	var ticket_default_group models.TicketDefaultGroup
	fmt.Print("Body", body)
	db.Where("id = ?", id).First(&ticket_default_group)

	// ticket_default_group.Name = body.Name

	db.Save(&ticket_default_group)
	return ticket_default_group
}