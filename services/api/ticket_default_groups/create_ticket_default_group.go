package api

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	"errors"
	validations "github.com/tejas-cogo/go-cogoport/services/validations"
)

type TicketDefaultGroupService struct {
	TicketDefaultGroup models.TicketDefaultGroup
}

func CreateTicketDefaultGroup(ticket_default_group models.TicketDefaultGroup) (models.TicketDefaultGroup,error) {
	db := config.GetDB()
	tx := db.Begin()
	var err error

	stmt := validations.ValidateTicketDefaultGroup(ticket_default_group)
	if stmt != "validated" {
		return ticket_default_group, errors.New(stmt)
	}

	ticket_default_group.Status = "active"

	if err := tx.Create(&ticket_default_group).Error; err != nil {
		tx.Rollback()
		return ticket_default_group, errors.New(err.Error())
	}

	tx.Commit()

	return ticket_default_group, err
}
