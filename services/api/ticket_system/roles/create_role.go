package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

type RoleService struct {
	Role models.Role
}

func CreateRole(role models.Role) models.Role {
	db := config.GetDB()
	// result := map[string]interface{}{}
	role.Status = "active"
	db.Create(&role)
	return role
}