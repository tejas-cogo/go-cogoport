package api

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	"errors"
	validations "github.com/tejas-cogo/go-cogoport/services/validations"
)

type TicketDefaultTypeService struct {
	TicketDefaultType models.TicketDefaultType
}

func CreateTicketDefaultType(ticket_default_type models.TicketDefaultType) (models.TicketDefaultType, error) {
	db := config.GetDB()
	tx := db.Begin()
	var err error

	stmt := validations.ValidateTicketDefaultType(ticket_default_type)
	if stmt != "validated" {
		return ticket_default_type, errors.New(stmt)
	}

	ticket_default_type.Status = "active"

	if err := tx.Create(&ticket_default_type).Error; err != nil {
		tx.Rollback()
		return ticket_default_type, errors.New("Cannot create ticket default type!")
	}

	tx.Commit()

	return ticket_default_type, err

}
