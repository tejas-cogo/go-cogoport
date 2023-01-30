package ticket_system

import (
	"github.com/tejas-cogo/go-cogoport/config"
	"github.com/tejas-cogo/go-cogoport/models"
)

func UpdateRole(body models.Role) (string, error, models.Role) {
	db := config.GetDB()
	tx := db.Begin()
	var err error

	var role models.Role

	if err := tx.Where("id = ?", body.ID).First(&role).Error; err != nil {
		tx.Rollback()
		return "Error Occurred!", err, body
	}
	
	if body.Name != "" {
		role.Name = body.Name
	}
	if body.Level != 0 {
		role.Level= body.Level
	}
	if body.Status != "" {
		role.Status = body.Status
	}

	if err := tx.Save(&role).Error; err != nil {
		tx.Rollback()
		return "Error Occurred!", err, body
	}

	tx.Commit()
	
	return "Successfully Updated!", err, role
}
