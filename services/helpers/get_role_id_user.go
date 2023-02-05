package apihelper

import (
	"github.com/gofrs/uuid"
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func GetRoleIdUser(RoleID uuid.UUID) uuid.UUID {
	db2 := config.GetCDB()
	// db := config.GetDB()
	var auth_role []models.AuthRole

	db2.Where("id = ? ", RoleID).First(&auth_role)

	return auth_role.StakeholderId

}
