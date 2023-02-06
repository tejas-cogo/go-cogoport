package api

import (
	"errors"

	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	// validations _"github.com/tejas-cogo/go-cogoport/services/validations"
)

type TicketDefaultRoleService struct {
	TicketDefaultRole models.TicketDefaultRole
}

func CreateTicketDefaultRole(ticket_default_role models.TicketDefaultRole) (models.TicketDefaultRole, error) {
	db := config.GetDB()
	tx := db.Begin()
	var err error
	var existed_default_role models.TicketDefaultRole

	// stmt := validations.ValidateTicketDefaultRole(ticket_default_role)
	// if stmt != "validated" {
	// 	return ticket_default_role, errors.New(stmt)
	// }

	ticket_default_role.Status = "active"

	tx.Where("ticket_default_type_id = ? and level = ? and status = ?", ticket_default_role.TicketDefaultTypeID, ticket_default_role.RoleID, ticket_default_role.UserID, ticket_default_role.Level, "active").First(&existed_default_role)

	if existed_default_role.ID > 0 {
		if err := tx.Model(&ticket_default_role).Where("id = ?", existed_default_role.ID).Update("status", "inactive").Error; err != nil {
			tx.Rollback()
			return ticket_default_role, errors.New(err.Error())
		}

		if err := tx.Where("id = ?", existed_default_role.ID).Delete(&ticket_default_role).Error; err != nil {
			tx.Rollback()
			return ticket_default_role, errors.New(err.Error())
		}
	}

	if err := tx.Create(&ticket_default_role).Error; err != nil {
		tx.Rollback()
		return ticket_default_role, errors.New(err.Error())
	}

	tx.Commit()

	return ticket_default_role, err
}
