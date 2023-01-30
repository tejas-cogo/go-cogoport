package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func UpdateTicketDefaultType(body models.TicketDefaultType) (string,error,models.TicketDefaultType) {
	db := config.GetDB()
	tx := db.Begin()
	var err error

	var ticket_default_type models.TicketDefaultType

	if err := tx.Where("id = ?", body.ID).Find(&ticket_default_type).Error; err != nil {
		tx.Rollback()
		return "Error Occurred!", err, body
	}

	if body.TicketType != "" {
		ticket_default_type.TicketType = body.TicketType
	}
	if body.AdditionalOptions != nil {
		ticket_default_type.AdditionalOptions = body.AdditionalOptions
	}
	if body.Status != "" {
		ticket_default_type.Status = body.Status
	}

	if err := tx.Save(&ticket_default_type).Error; err != nil {
		tx.Rollback()
		return "Error Occurred!", err, body
	}

	tx.Commit()

	return "Sucessfully Updated!", err, ticket_default_type
}
