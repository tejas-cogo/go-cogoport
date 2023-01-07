package ticket_system

import (
	"fmt"
	"github.com/ChandelShikha/go-cogoport/config"
	"github.com/ChandelShikha/go-cogoport/models"
)

func UpdateRole(id uint, body models.Role) models.Role {
	db := config.GetDB()
	var role models.Role
	fmt.Print("Body", body)
	db.Where("id = ?", id).First(&role)

	// role.Name = body.Name

	db.Save(&role)
	return role
}