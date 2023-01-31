package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	"errors"
)

type TicketDefaultGroupService struct {
	TicketDefaultGroup models.TicketDefaultGroup
}

func CreateTicketDefaultGroup(ticket_default_group models.TicketDefaultGroup) (models.TicketDefaultGroup,error) {
	db := config.GetDB()
	tx := db.Begin()
	var err error

	stmt := validate(ticket_default_group)
	if stmt != "validated" {
		return ticket_default_group, errors.New(stmt)
	}

	ticket_default_group.Status = "active"

	if err := tx.Create(&ticket_default_group).Error; err != nil {
		tx.Rollback()
		return ticket_default_group, errors.New("Error Occurred!")
	}

	tx.Commit()

	return ticket_default_group, err
}

func validate(ticket_default_group models.TicketDefaultGroup) string {

	if ticket_default_group.GroupID == 0 {
		return ("Group Is Required!")
	}

	if ticket_default_group.TicketDefaultTypeID == 0 {
		return ("TicketDefaultTypeID Is Required!")
	}

	return ("validated")
}
