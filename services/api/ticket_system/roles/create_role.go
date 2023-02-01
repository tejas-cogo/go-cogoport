package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	"errors"
	validations "github.com/tejas-cogo/go-cogoport/services/validations"
)

type RoleService struct {
	Role models.Role
}

func CreateRole(role models.Role) (models.Role,error) {
	db := config.GetDB()
	tx := db.Begin()
	var err error

	stmt := validations.validate_role(role)
	if stmt != "validated" {
		return role, errors.New(stmt)
	}

	role.Status = "active"

	if err := tx.Create(&role).Error; err != nil {
		tx.Rollback()
		return role, errors.New("Error Occurred!")
	}

	tx.Commit()

	return role, err
}
