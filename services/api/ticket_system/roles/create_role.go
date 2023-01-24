package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

type RoleService struct {
	Role models.Role
}

func CreateRole(role models.Role) (string, error) {
	db := config.GetDB()
	tx := db.Begin()
	var err error

	stmt := validate(role)
	if stmt != "validated" {
		return stmt, err
	}

	role.Status = "active"

	if err := tx.Create(&role).Error; err != nil {
		tx.Rollback()
		return "Error Occurred", err
	}

	tx.Commit()

	return "Successfully Created", err
}

func validate(role models.Role) string {
	if role.Name == "" {
		return ("Role Name Is Required")
	}

	if role.Level == 0 {
		return ("Level Is Required")
	}

	return ("validated")
}
