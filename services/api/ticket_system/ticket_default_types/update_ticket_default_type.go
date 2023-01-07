package ticket_system

import (
	"fmt"
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func UpdateTicketDefaultType(id uint, body models.TicketDefaultType) models.TicketDefaultType {
	db := config.GetDB()
	var ticket_default_type models.TicketDefaultType
	fmt.Print("Body", body)
	db.Where("id = ?", id).First(&ticket_default_type)

	// ticket_default_type.Name = body.Name

	db.Save(&ticket_default_type)
	return ticket_default_type
}