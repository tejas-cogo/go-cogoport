package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func UpdateTicketDefaultType(body models.TicketDefaultType) models.TicketDefaultType {
	db := config.GetDB()
	var ticket_default_type models.TicketDefaultType
	db.Where("id = ?", body.ID).Find(&ticket_default_type)

	if body.TicketType != "" {
		ticket_default_type.TicketType = body.TicketType
	}
	if body.AdditionalOptions != nil {
		ticket_default_type.AdditionalOptions = body.AdditionalOptions
	}
	if body.Status != "" {
		ticket_default_type.Status = body.Status
	}

	db.Save(&ticket_default_type)
	return ticket_default_type
}
