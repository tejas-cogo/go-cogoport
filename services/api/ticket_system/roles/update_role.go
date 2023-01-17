package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func UpdateRole(body models.Role) models.Role {
	db := config.GetDB()
	var role models.Role
	db.Where("id = ?", body.ID).First(&role)

	if body.Name != "" {
		role.Name = body.Name
	}
	if body.Level != 0 {
		role.Level= body.Level
	}
	if body.Status != "" {
		role.Status = body.Status
	}

	db.Save(&role)
	return role
}
