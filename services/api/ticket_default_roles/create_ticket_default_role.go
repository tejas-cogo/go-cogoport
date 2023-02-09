package api

import (
	"errors"
	"fmt"

	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	validations "github.com/tejas-cogo/go-cogoport/services/validations"
)

type TicketDefaultRoleService struct {
	TicketDefaultRole models.TicketDefaultRole
}

func CreateTicketDefaultRole(ticket_default_role models.TicketDefaultRole) (models.TicketDefaultRole, error) {
	db := config.GetDB()
	tx := db.Begin()
	var err error
	var existed_default_role models.TicketDefaultRole

	tx.Where("ticket_default_type_id = ? and level = ? and status = ?", ticket_default_role.TicketDefaultTypeID, ticket_default_role.Level, "active").First(&existed_default_role)

	if existed_default_role.ID > 0 {
		if err := tx.Model(&ticket_default_role).Where("id = ?", existed_default_role.ID).Update("status", "inactive").Delete(&existed_default_role).Error; err != nil {
			tx.Rollback()
			return ticket_default_role, errors.New(err.Error())
		}

		// if err := tx.Where("id = ?", existed_default_role.ID).Delete(&ticket_default_role).Error; err != nil {
		// 	tx.Rollback()
		// 	return ticket_default_role, errors.New(err.Error())
		// }
	}

	ticket_default_role.Status = "active"

	stmt1 := validations.ValidateTicketDefaultRole(ticket_default_role)
	if stmt1 != "validated" {
		return ticket_default_role, errors.New(stmt1)
	}

	stmt2 := validations.ValidateDuplicateDefaultType(ticket_default_role)
	if stmt2 != "validated" {
		return ticket_default_role, errors.New(stmt2)
	}
	if err := tx.Create(&ticket_default_role).Error; err != nil {
		tx.Rollback()
		return ticket_default_role, errors.New(err.Error())
	}

	tx.Commit()
	fmt.Println("df", ticket_default_role.Status)

	return ticket_default_role, err
}
