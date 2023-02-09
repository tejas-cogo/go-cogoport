package api

import (
	"errors"

	"github.com/google/uuid"
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	validations "github.com/tejas-cogo/go-cogoport/services/validations"
)

func UpdateTicketDefaultRole(body models.TicketDefaultRole) (models.TicketDefaultRole, error) {
	db := config.GetDB()
	tx := db.Begin()
	var err error

	var ticket_default_role models.TicketDefaultRole

	if err := tx.Where("id = ?", body.ID).Find(&ticket_default_role).Error; err != nil {
		tx.Rollback()
		return body, errors.New(err.Error())
	}

	if body.TicketDefaultTypeID > 0 {
		ticket_default_role.TicketDefaultTypeID = body.TicketDefaultTypeID
	}
	if body.RoleID != uuid.Nil {
		ticket_default_role.RoleID = body.RoleID
	}
	if body.UserID != uuid.Nil {
		ticket_default_role.UserID = body.UserID
	}
	if body.Status != "" {
		ticket_default_role.Status = body.Status
	}

	stmt2 := validations.ValidateDuplicateDefaultType(ticket_default_role)
	if stmt2 != "validated" {
		return ticket_default_role, errors.New(stmt2)
	}

	if err := tx.Save(&ticket_default_role).Error; err != nil {
		tx.Rollback()
		return body, errors.New(err.Error())
	}

	tx.Commit()

	return ticket_default_role, err
}
