package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func ListRole() []models.Role {
	db := config.GetDB()

	var role []models.Role

	db.Find(&role)

	return role
}
