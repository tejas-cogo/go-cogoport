package ticket_system

import (
	"github.com/ChandelShikha/go-cogoport/config"
	"github.com/ChandelShikha/go-cogoport/models"
)

type RoleService struct {
	Role models.Role
}

func CreateRole(role models.Role) models.Role {
	db := config.GetDB()
	// result := map[string]interface{}{}
	db.Create(&role)
	return role
}