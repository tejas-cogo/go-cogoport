package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
	"errors"
)

type RoleService struct {
	Role models.Role
}

func CreateRole(role models.Role) (models.Role,error) {
	db := config.GetDB()
	tx := db.Begin()
	var err error

	stmt := validate(role)
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

func validate(role models.Role) string {
	if role.Name == "" {
		return ("Role Name Is Required!")
	}

	if role.Level == 0 {
		return ("Level Is Required!")
	}

	if role.Level > 9 {
		return ("Level should be in range 1-9!")
	}

	if len(role.Name) < 2 || len(role.Name) > 30 {
		return ("Role field must be between 2-30 chars!")
	}

	return ("validated")
}
