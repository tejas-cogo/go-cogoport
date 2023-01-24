package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

type TicketDefaultTypeService struct {
	TicketDefaultType models.TicketDefaultType
}

func CreateTicketDefaultType(ticket_default_type models.TicketDefaultType) (string,error) {
	db := config.GetDB()
	tx := db.Begin()
	var err error

	stmt := validate(ticket_default_type)
	if stmt != "validated" {
		return stmt, err
	}

	ticket_default_type.Status = "active"

	if err := tx.Create(&ticket_default_type).Error; err != nil {
		tx.Rollback()
		return "Error Occurred", err
	}

	tx.Commit()

	return "Successfully Created", err

}

func validate(ticket_default_type models.TicketDefaultType) string {
	if ticket_default_type.TicketType == "" {
		return ("TicketType Is Required")
	}

	return ("validated")
}
