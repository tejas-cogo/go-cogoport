package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	"errors"
)

type TicketDefaultTypeService struct {
	TicketDefaultType models.TicketDefaultType
}

func CreateTicketDefaultType(ticket_default_type models.TicketDefaultType) (models.TicketDefaultType, error) {
	db := config.GetDB()
	tx := db.Begin()
	var err error

	stmt := validate(ticket_default_type)
	if stmt != "validated" {
		return ticket_default_type, errors.New(stmt)
	}

	ticket_default_type.Status = "active"

	if err := tx.Create(&ticket_default_type).Error; err != nil {
		tx.Rollback()
		return ticket_default_type, errors.New("Error Occurred!")
	}

	tx.Commit()

	return ticket_default_type, err

}

func validate(ticket_default_type models.TicketDefaultType) string {
	if ticket_default_type.TicketType == "" {
		return ("TicketType Is Required!")
	}

	return ("validated")
}
