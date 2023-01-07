package ticket_system

import (
	"github.com/ChandelShikha/go-cogoport/config"
	"github.com/ChandelShikha/go-cogoport/models"
)

func ListRole() []models.Role {
	db := config.GetDB()

	var role []models.Role

	result := map[string]interface{}{}
	db.Find(&role).Take(&result)

	return role
}
