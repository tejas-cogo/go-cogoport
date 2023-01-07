package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func ListRole() []models.Role {
	db := config.GetDB()

	var role []models.Role

	result := map[string]interface{}{}
	db.Find(&role).Take(&result)

	return role
}
