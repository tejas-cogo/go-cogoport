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
		return "Error Occurred!", err
	}

	tx.Commit()

	return "Successfully Created!", err
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
